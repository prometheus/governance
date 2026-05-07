// collect-write-access: Collect GitHub users with write access across
// Prometheus orgs, with full provenance.
//
// Lists org members and per-repo direct collaborators with push permission.
// Requires admin access on each org/repo for the collaborators listing.
//
// Usage:
//
//	collect-write-access [-out <dir>]
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
	"sort"
	"strings"
	"time"

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

// Writer is one (login, source, access_type) triple.
type Writer struct {
	Login      string
	Source     string
	AccessType string
}

func main() {
	outDir := flag.String("out", "out-access", "output directory")
	flag.Parse()

	if err := run(*outDir); err != nil {
		logf("ERROR: %v", err)
		os.Exit(1)
	}
}

func run(outDir string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create out dir: %w", err)
	}

	token, err := resolveToken()
	if err != nil {
		return fmt.Errorf("resolve token: %w", err)
	}

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	generatedAt := time.Now().UTC().Format(time.RFC3339)

	var (
		writers  []Writer
		manifest []ManifestEntry
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

	// Direct collaborators per repo.
	for _, org := range orgs {
		logf("Enumerating repos in %s for direct collaborators …", org)
		repos, err := listNonArchivedRepos(ctx, client, org)
		if err != nil {
			logf("SKIP %s — could not list repos: %v", org, err)
			continue
		}

		for _, repo := range repos {
			collabs, err := listPushCollaborators(ctx, client, org, repo)
			if err != nil {
				logf("SKIP %s/%s — %v", org, repo, err)
				continue
			}
			if len(collabs) == 0 {
				continue
			}
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
	}

	// Write manifest.
	if err := writeManifest(filepath.Join(outDir, "manifest.json"), manifest); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}

	// Write raw CSV.
	rawPath := filepath.Join(outDir, "write-access-raw.csv")
	rawRows := make([][]string, len(writers))
	for i, wr := range writers {
		rawRows[i] = []string{wr.Login, wr.Source, wr.AccessType}
	}
	if err := writeCSV(rawPath, generatedAt, []string{"login", "source", "access_type"}, rawRows); err != nil {
		return fmt.Errorf("write raw csv: %w", err)
	}

	// Deduplicate, sort, write final CSV.
	dedup := deduplicate(writers)
	finalPath := filepath.Join(outDir, "write-access.csv")
	dedupRows := make([][]string, len(dedup))
	for i, e := range dedup {
		dedupRows[i] = []string{e.Login, strings.Join(e.Sources, ";"), strings.Join(e.AccessTypes, ";")}
	}
	if err := writeCSV(finalPath, generatedAt, []string{"login", "sources", "access_types"}, dedupRows); err != nil {
		return fmt.Errorf("write final csv: %w", err)
	}

	logf("Done.")
	logf("  Generated at : %s", generatedAt)
	logf("  Raw entries  : %d -> %s", len(writers), rawPath)
	logf("  Deduplicated : %d -> %s", len(dedup), finalPath)
	logf("  Manifest     : %s", filepath.Join(outDir, "manifest.json"))

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

// DedupEntry is one merged voter row.
type DedupEntry struct {
	Login       string
	Sources     []string
	AccessTypes []string
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

// writeCSV writes a CSV file with a "# Generated at:" comment, a header row,
// and the given data rows.
func writeCSV(path, generatedAt string, header []string, rows [][]string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	if _, err := fmt.Fprintf(f, "# Generated at: %s\n", generatedAt); err != nil {
		return err
	}

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

