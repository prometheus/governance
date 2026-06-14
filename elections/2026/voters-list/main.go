// collect-write-access: Collect GitHub users with write access across
// Prometheus orgs, with full provenance.
//
// Lists org members and per-repo direct collaborators with push permission.
// Requires admin access on each org/repo for the collaborators listing.
//
// It also reads each repo's root MAINTAINERS(.md) file to capture maintainer
// names and emails: these enrich the write-access list where a maintainer's
// GitHub handle matches a known login, and maintainers not in the write-access
// set are written to maintainers-without-write-access.csv (they do not become
// voters).
//
// A mapping file (-mapping, default email-to-github.csv) ties a maintainer email
// to a GitHub login and is applied in both directions. An email resolves (or
// corrects) the GitHub handle, so a MAINTAINERS entry listed by name+email only
// (no @handle) still matches the write-access set; a handle overrides the email,
// so the list is the authoritative source for a maintainer's contact address.
//
// Usage:
//
//	collect-write-access [-out <dir>] [-mapping <file>]
//
// Auth: reads GITHUB_TOKEN env var, or falls back to `gh auth token`.
package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/google/go-github/v66/github"
)

var (
	orgs        = []string{"prometheus", "prometheus-community"}
	excludeBots = map[string]bool{
		"prombot":              true,
		"prom-label-proxy-bot": true,
	}
)

// ManifestEntry records one API call's provenance.
type ManifestEntry struct {
	Org       string `json:"org"`
	Repo      string `json:"repo,omitempty"`
	Type      string `json:"type"`
	Count     int    `json:"count"`
	FetchedAt string `json:"fetched_at"`
}

// Writer is one (login, source, access_type) triple. Name and Email are only
// populated for rows sourced from a MAINTAINERS file (access_type "maintainer").
type Writer struct {
	Login      string
	Source     string
	AccessType string
	Name       string
	Email      string
}

// Maintainer is one entry parsed from a repo's root MAINTAINERS(.md) file.
type Maintainer struct {
	Org    string
	Repo   string
	Name   string
	Email  string
	Handle string // GitHub login, without the leading "@"; may be empty.
}

func main() {
	outDir := flag.String("out", "out-access", "output directory")
	mappingPath := flag.String("mapping", "email-to-github.csv", "email->github login mapping file")
	flag.Parse()

	if err := run(*outDir, *mappingPath); err != nil {
		logf("ERROR: %v", err)
		os.Exit(1)
	}
}

func run(outDir, mappingPath string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create out dir: %w", err)
	}

	emailToLogin, loginToEmail, err := loadMapping(mappingPath)
	if err != nil {
		return fmt.Errorf("load mapping: %w", err)
	}
	logf("Loaded %d email<->github mappings from %s", len(emailToLogin), mappingPath)

	token, err := resolveToken()
	if err != nil {
		return fmt.Errorf("resolve token: %w", err)
	}

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	generatedAt := time.Now().UTC().Format(time.RFC3339)

	var (
		writers     []Writer
		maintainers []Maintainer
		manifest    []ManifestEntry
	)

	// Org members.
	for _, org := range orgs {
		logf("Fetching org members for %s …", org)
		members, err := listOrgMembers(ctx, client, org)
		if err != nil {
			logf("SKIP %s — %v", org, err)
			continue
		}
		for _, m := range members {
			writers = append(writers, Writer{
				Login:      m,
				Source:     org,
				AccessType: "org_member",
			})
		}
		manifest = append(manifest, ManifestEntry{
			Org:       org,
			Type:      "org_members",
			Count:     len(members),
			FetchedAt: time.Now().UTC().Format(time.RFC3339),
		})
		logf("OK   %s — %d org members", org, len(members))
	}

	// Direct collaborators and MAINTAINERS files per repo.
	for _, org := range orgs {
		logf("Enumerating repos in %s for direct collaborators and maintainers …", org)
		repos, err := listNonArchivedRepos(ctx, client, org)
		if err != nil {
			logf("SKIP %s — could not list repos: %v", org, err)
			continue
		}

		for _, repo := range repos {
			collabs, err := listPushCollaborators(ctx, client, org, repo)
			if err != nil {
				logf("SKIP %s/%s — collaborators: %v", org, repo, err)
			} else if len(collabs) > 0 {
				for _, c := range collabs {
					writers = append(writers, Writer{
						Login:      c,
						Source:     org + "/" + repo,
						AccessType: "direct_collaborator",
					})
				}
				manifest = append(manifest, ManifestEntry{
					Org:       org,
					Repo:      repo,
					Type:      "direct_collaborators",
					Count:     len(collabs),
					FetchedAt: time.Now().UTC().Format(time.RFC3339),
				})
				logf("OK   %s/%s — %d direct collaborators with push", org, repo, len(collabs))
			}

			ms, found, err := fetchMaintainers(ctx, client, org, repo)
			if err != nil {
				logf("SKIP %s/%s — maintainers: %v", org, repo, err)
			} else if found {
				maintainers = append(maintainers, ms...)
				manifest = append(manifest, ManifestEntry{
					Org:       org,
					Repo:      repo,
					Type:      "maintainers_file",
					Count:     len(ms),
					FetchedAt: time.Now().UTC().Format(time.RFC3339),
				})
				logf("OK   %s/%s — %d maintainers parsed", org, repo, len(ms))
			}
		}
	}

	// Classify maintainers against the write-access set (matched by handle,
	// case-insensitively). Matched entries enrich existing voters with
	// name/email; unmatched entries (or those with no handle) are reported
	// separately and do not become voters. The mapping is applied in both
	// directions and takes precedence over the MAINTAINERS file: an email
	// resolves (or corrects) the handle, fixing entries listed by name+email
	// only, and a handle overrides the email with the list's address.
	loginByLower := make(map[string]string, len(writers))
	for _, w := range writers {
		loginByLower[strings.ToLower(w.Login)] = w.Login
	}
	var noAccess []Maintainer
	for _, m := range maintainers {
		if id, ok := emailToLogin[strings.ToLower(m.Email)]; ok && m.Email != "" {
			m.Handle = id
		}
		if email, ok := loginToEmail[strings.ToLower(m.Handle)]; ok && m.Handle != "" {
			m.Email = email
		}
		canon, ok := loginByLower[strings.ToLower(m.Handle)]
		if m.Handle != "" && ok {
			writers = append(writers, Writer{
				Login:      canon,
				Source:     m.Org + "/" + m.Repo,
				AccessType: "maintainer",
				Name:       m.Name,
				Email:      m.Email,
			})
			continue
		}
		noAccess = append(noAccess, m)
	}

	// The mapping is authoritative for contact addresses: attach (or override)
	// the email of any voter whose login matches a github_id in the mapping.
	// This also reaches voters sourced only from org membership or direct
	// collaboration, who have no MAINTAINERS entry to carry an email.
	for i := range writers {
		if email, ok := loginToEmail[strings.ToLower(writers[i].Login)]; ok {
			writers[i].Email = email
		}
	}

	// Write manifest.
	if err := writeManifest(filepath.Join(outDir, "manifest.json"), manifest); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}

	// Write raw CSV.
	rawPath := filepath.Join(outDir, "write-access-raw.csv")
	rawRows := make([][]string, len(writers))
	for i, wr := range writers {
		rawRows[i] = []string{wr.Login, wr.Source, wr.AccessType, wr.Name, wr.Email}
	}
	if err := writeCSV(rawPath, []string{"login", "source", "access_type", "name", "email"}, rawRows); err != nil {
		return fmt.Errorf("write raw csv: %w", err)
	}

	// Deduplicate, sort, write final CSV.
	dedup := deduplicate(writers)
	finalPath := filepath.Join(outDir, "write-access.csv")
	dedupRows := make([][]string, len(dedup))
	for i, e := range dedup {
		dedupRows[i] = []string{
			e.Login,
			strings.Join(e.Sources, ";"),
			strings.Join(e.AccessTypes, ";"),
			strings.Join(e.Names, ";"),
			strings.Join(e.Emails, ";"),
		}
	}
	if err := writeCSV(finalPath, []string{"login", "sources", "access_types", "name", "email"}, dedupRows); err != nil {
		return fmt.Errorf("write final csv: %w", err)
	}

	// Write maintainers that are not in the write-access set.
	noAccessPath := filepath.Join(outDir, "maintainers-without-write-access.csv")
	noAccessRows := dedupMaintainers(noAccess)
	if err := writeCSV(noAccessPath, []string{"name", "email", "github_handle", "repos"}, noAccessRows); err != nil {
		return fmt.Errorf("write maintainers-without-write-access csv: %w", err)
	}

	logf("Done.")
	logf("  Generated at      : %s", generatedAt)
	logf("  Raw entries       : %d -> %s", len(writers), rawPath)
	logf("  Deduplicated      : %d -> %s", len(dedup), finalPath)
	logf("  Maintainers w/o WA : %d -> %s", len(noAccessRows), noAccessPath)
	logf("  Manifest          : %s", filepath.Join(outDir, "manifest.json"))

	return nil
}

func resolveToken() (string, error) {
	if t := os.Getenv("GITHUB_TOKEN"); t != "" {
		return t, nil
	}
	out, err := exec.Command("gh", "auth", "token").Output()
	if err != nil {
		return "", errors.New("no GITHUB_TOKEN env var and `gh auth token` failed; run `gh auth login` or set GITHUB_TOKEN")
	}
	return strings.TrimSpace(string(out)), nil
}

// loadMapping reads the maintainer email <-> GitHub-login mapping from a CSV
// file with "email,github_id" rows ('#' comment lines are ignored). Each row is
// an identity assertion usable in both directions, so two maps are returned:
// emailToLogin (key lower-cased) and the reverse loginToEmail (key lower-cased,
// value the email as written). A missing file is not an error: it returns empty
// maps.
func loadMapping(path string) (emailToLogin, loginToEmail map[string]string, err error) {
	emailToLogin = map[string]string{}
	loginToEmail = map[string]string{}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			logf("mapping file %s not found; continuing without it", path)
			return emailToLogin, loginToEmail, nil
		}
		return nil, nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comment = '#'
	r.FieldsPerRecord = -1
	rows, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	for _, row := range rows {
		if len(row) < 2 {
			continue
		}
		email := strings.TrimSpace(row[0])
		id := strings.TrimSpace(row[1])
		if email == "" || id == "" || strings.EqualFold(email, "email") {
			continue
		}
		emailToLogin[strings.ToLower(email)] = id
		loginToEmail[strings.ToLower(id)] = email
	}
	return emailToLogin, loginToEmail, nil
}

func listOrgMembers(ctx context.Context, client *github.Client, org string) ([]string, error) {
	var all []string
	opt := &github.ListMembersOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		users, resp, err := client.Organizations.ListMembers(ctx, org, opt)
		if err != nil {
			return nil, err
		}
		for _, u := range users {
			all = append(all, u.GetLogin())
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return all, nil
}

func listNonArchivedRepos(ctx context.Context, client *github.Client, org string) ([]string, error) {
	var all []string
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return nil, err
		}
		for _, r := range repos {
			if r.GetArchived() {
				continue
			}
			all = append(all, r.GetName())
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return all, nil
}

func listPushCollaborators(ctx context.Context, client *github.Client, org, repo string) ([]string, error) {
	var all []string
	opt := &github.ListCollaboratorsOptions{
		Affiliation: "direct",
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		users, resp, err := client.Repositories.ListCollaborators(ctx, org, repo, opt)
		if err != nil {
			return nil, err
		}
		for _, u := range users {
			if u.GetPermissions()["push"] {
				all = append(all, u.GetLogin())
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return all, nil
}

// maintainerFileNames are the root paths we look for, in priority order.
var maintainerFileNames = []string{"MAINTAINERS.md", "MAINTAINER.md"}

// fetchMaintainers fetches and parses a repo's root MAINTAINERS file. found is
// false (with nil error) when no such file exists.
func fetchMaintainers(ctx context.Context, client *github.Client, org, repo string) (entries []Maintainer, found bool, err error) {
	for _, name := range maintainerFileNames {
		fc, _, resp, err := client.Repositories.GetContents(ctx, org, repo, name, nil)
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				continue
			}
			return nil, false, err
		}
		content, err := fc.GetContent()
		if err != nil {
			return nil, false, err
		}
		parsed := parseMaintainers(content)
		for i := range parsed {
			parsed[i].Org = org
			parsed[i].Repo = repo
		}
		return parsed, true, nil
	}
	return nil, false, nil
}

var (
	emailRE  = regexp.MustCompile(`[\w.+%-]+@[\w-]+\.[\w.-]+`)
	handleRE = regexp.MustCompile(`@([A-Za-z0-9-]+)`)
	// githubURLRE captures the handle from a GitHub profile link, e.g.
	// "[Ben Reedy](https://github.com/breed808)". The trailing delimiter keeps
	// it from matching repo links like "github.com/org/repo/blob/…".
	githubURLRE = regexp.MustCompile(`github\.com/([A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?)(?:[)>?#\s]|$)`)
	// mdLinkRE matches a markdown link; replacing with $1 keeps the link text.
	mdLinkRE = regexp.MustCompile(`\[([^\]]*)\]\([^)]*\)`)
	// parenHandleRE matches a lone parenthesised token used as a handle, e.g.
	// "(breed808)". Excludes anything containing @ or . so emails don't match.
	parenHandleRE = regexp.MustCompile(`\(([A-Za-z0-9][A-Za-z0-9-]*)\)`)
	// listMarkerRE strips leading bullet markers and an optional `component`:
	// prefix, e.g. "* `tsdb`: ".
	listMarkerRE = regexp.MustCompile("^[\\s>]*[*•-]?\\s*(?:`[^`]*`\\s*:?\\s*)*")
	alumniRE     = regexp.MustCompile(`(?i)alumni|emeritus|former|past|previous|inactive|retired`)
)

// parseMaintainers extracts maintainer entries from a MAINTAINERS file. It is
// deliberately lenient and best-effort: the ecosystem uses several formats
// (`Name <email> @handle`, `Name (email / @handle)`, `Name (handle) - email`,
// component-scoped and multi-maintainer lines). Prose intro lines may parse
// imperfectly, but such people are typically org members (hence voters) anyway,
// so only their name/email enrichment is affected.
func parseMaintainers(content string) []Maintainer {
	var out []Maintainer
	skip := false // true while inside an alumni/former-maintainer section.

	for _, rawLine := range strings.Split(content, "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}

		if heading, ok := sectionHeading(line); ok {
			skip = alumniRE.MatchString(heading)
			continue
		}
		if skip {
			continue
		}

		body := listMarkerRE.ReplaceAllString(line, "")
		// Names never contain commas, so commas separate maintainers on a line.
		for _, chunk := range strings.Split(body, ",") {
			if m, ok := parseMaintainerChunk(chunk); ok {
				out = append(out, m)
			}
		}
	}
	return out
}

// sectionHeading returns the heading text if line looks like a section heading
// (a markdown "#" heading, a wholly-bold line, or a short colon-terminated
// label with no email), so callers can detect alumni sections.
func sectionHeading(line string) (string, bool) {
	switch {
	case strings.HasPrefix(line, "#"):
		return strings.TrimLeft(line, "# "), true
	case strings.HasPrefix(line, "**") && strings.HasSuffix(line, "**"):
		return strings.Trim(line, "*: "), true
	case strings.HasPrefix(line, "__") && strings.HasSuffix(line, "__"):
		return strings.Trim(line, "_: "), true
	case strings.HasSuffix(line, ":") && !emailRE.MatchString(line) && !strings.ContainsAny(line, "*-•"):
		return strings.TrimRight(line, ": "), true
	case isPlainHeading(line):
		// A short capitalised line with no contact info, e.g. a bare "Alumni".
		return line, true
	}
	return "", false
}

// isPlainHeading reports whether line is a short, capitalised label carrying no
// maintainer contact info (email, handle, or profile link) — i.e. a section
// header such as "Alumni" or "Maintainers in alphabetical order".
func isPlainHeading(line string) bool {
	r := []rune(line)
	if len(r) == 0 || !unicode.IsUpper(r[0]) {
		return false
	}
	if emailRE.MatchString(line) || strings.Contains(line, "@") ||
		strings.Contains(line, "](") || strings.Contains(line, "github.com") {
		return false
	}
	return len(strings.Fields(line)) <= 4
}

// parseMaintainerChunk parses a single maintainer from one comma-delimited chunk.
func parseMaintainerChunk(chunk string) (Maintainer, bool) {
	var m Maintainer

	if loc := emailRE.FindStringIndex(chunk); loc != nil {
		m.Email = chunk[loc[0]:loc[1]]
	}
	// Strip emails first so the "@" in an address isn't mistaken for a handle.
	noEmail := emailRE.ReplaceAllString(chunk, " ")
	switch {
	case handleRE.MatchString(noEmail):
		m.Handle = handleRE.FindStringSubmatch(noEmail)[1]
	case githubURLRE.MatchString(chunk):
		// Markdown profile link, e.g. "[Name](https://github.com/handle)".
		m.Handle = githubURLRE.FindStringSubmatch(chunk)[1]
	case m.Email != "" && parenHandleRE.MatchString(noEmail):
		// A bare "(token)" is only treated as a handle alongside an email
		// (e.g. "Name (breed808) - email"); otherwise prose like
		// "maintainer(s)" would yield bogus handles.
		m.Handle = parenHandleRE.FindStringSubmatch(noEmail)[1]
	}

	// Name is the leading text before the first email/handle/delimiter. Collapse
	// markdown links to their text ("[Ben Reedy](url)" -> "Ben Reedy"), then cut
	// at the email first so its local part isn't kept when it directly follows
	// the name with no bracket (e.g. "Amy Super amy.super@grafana.com").
	name := mdLinkRE.ReplaceAllString(chunk, "$1")
	if loc := emailRE.FindStringIndex(name); loc != nil {
		name = name[:loc[0]]
	}
	for _, sep := range []string{"<", "(", "@"} {
		if i := strings.Index(name, sep); i >= 0 {
			name = name[:i]
		}
	}
	// Drop any leading "label:" prose left over from component-scoped lines
	// (e.g. "`Makefile` and related build configuration: Simon Pasquier").
	if i := strings.LastIndex(name, ":"); i >= 0 {
		name = name[i+1:]
	}
	m.Name = strings.Trim(name, " \t/-:[]")

	if m.Email == "" && m.Handle == "" {
		return Maintainer{}, false
	}
	return m, true
}

// DedupEntry is one merged voter row.
type DedupEntry struct {
	Login       string
	Sources     []string
	AccessTypes []string
	Names       []string
	Emails      []string
}

func deduplicate(writers []Writer) []DedupEntry {
	byLogin := make(map[string]*DedupEntry)
	for _, w := range writers {
		if excludeBots[strings.ToLower(w.Login)] {
			continue
		}
		e, ok := byLogin[w.Login]
		if !ok {
			e = &DedupEntry{Login: w.Login}
			byLogin[w.Login] = e
		}
		e.Sources = append(e.Sources, w.Source)
		e.AccessTypes = append(e.AccessTypes, w.AccessType)
		e.Names = appendUnique(e.Names, w.Name)
		e.Emails = appendUnique(e.Emails, w.Email)
	}

	out := make([]DedupEntry, 0, len(byLogin))
	for _, e := range byLogin {
		out = append(out, *e)
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.ToLower(out[i].Login) < strings.ToLower(out[j].Login)
	})
	return out
}

// appendUnique appends v to s if v is non-empty and not already present.
func appendUnique(s []string, v string) []string {
	if v == "" {
		return s
	}
	for _, e := range s {
		if e == v {
			return s
		}
	}
	return append(s, v)
}

// dedupMaintainers merges maintainer entries that refer to the same person
// (keyed by handle, else email, else name — case-insensitively) and returns
// CSV rows of [name, email, github_handle, repos], sorted by name.
func dedupMaintainers(ms []Maintainer) [][]string {
	type agg struct {
		name, email, handle string
		repos               []string
	}
	byKey := make(map[string]*agg)
	for _, m := range ms {
		key := strings.ToLower(firstNonEmpty(m.Handle, m.Email, m.Name))
		if key == "" {
			continue
		}
		a, ok := byKey[key]
		if !ok {
			a = &agg{name: m.Name, email: m.Email, handle: m.Handle}
			byKey[key] = a
		}
		if a.name == "" {
			a.name = m.Name
		}
		if a.email == "" {
			a.email = m.Email
		}
		if a.handle == "" {
			a.handle = m.Handle
		}
		a.repos = appendUnique(a.repos, m.Org+"/"+m.Repo)
	}

	rows := make([][]string, 0, len(byKey))
	for _, a := range byKey {
		sort.Strings(a.repos)
		rows = append(rows, []string{a.name, a.email, a.handle, strings.Join(a.repos, ";")})
	}
	sort.Slice(rows, func(i, j int) bool {
		if li, lj := strings.ToLower(rows[i][0]), strings.ToLower(rows[j][0]); li != lj {
			return li < lj
		}
		return rows[i][1] < rows[j][1]
	})
	return rows
}

// firstNonEmpty returns the first non-empty string among its arguments.
func firstNonEmpty(vs ...string) string {
	for _, v := range vs {
		if v != "" {
			return v
		}
	}
	return ""
}

func writeManifest(path string, entries []ManifestEntry) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(entries)
}

// writeCSV writes a CSV file with a header row and the given data rows. The
// generation timestamp is recorded in manifest.json rather than as a CSV
// comment, so the output stays valid CSV (a leading comment line confuses
// strict parsers that expect every row to match the header's column count).
func writeCSV(path string, header []string, rows [][]string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	w := csv.NewWriter(f)

	if err := w.Write(header); err != nil {
		return err
	}
	for _, row := range rows {
		if err := w.Write(row); err != nil {
			return err
		}
	}

	w.Flush()
	return w.Error()
}

func logf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "[%s] %s\n", time.Now().UTC().Format(time.RFC3339), fmt.Sprintf(format, args...))
}
