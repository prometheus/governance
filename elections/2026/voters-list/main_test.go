package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseMaintainers(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []Maintainer
	}{
		{
			name:    "angle-bracket email and handle",
			content: "* Ben Kochie <superq@gmail.com> @SuperQ\n* Johannes 'fish' Ziemke <github@freigeist.org> @discordianfish\n",
			want: []Maintainer{
				{Name: "Ben Kochie", Email: "superq@gmail.com", Handle: "SuperQ"},
				{Name: "Johannes 'fish' Ziemke", Email: "github@freigeist.org", Handle: "discordianfish"},
			},
		},
		{
			name:    "email and handle in parens",
			content: "- André Bauer (monotek23@gmail.com / @monotek)\n- Scott Rigby (scott@r6by.com / @scottrigby)\n",
			want: []Maintainer{
				{Name: "André Bauer", Email: "monotek23@gmail.com", Handle: "monotek"},
				{Name: "Scott Rigby", Email: "scott@r6by.com", Handle: "scottrigby"},
			},
		},
		{
			name:    "handle in parens without @, email after dash",
			content: "- Ben Reedy (breed808) - breed808@breed808.com\n- Jan-Otto Kröpke (jkroepke) - github@jkroepke.de\n",
			want: []Maintainer{
				{Name: "Ben Reedy", Email: "breed808@breed808.com", Handle: "breed808"},
				{Name: "Jan-Otto Kröpke", Email: "github@jkroepke.de", Handle: "jkroepke"},
			},
		},
		{
			name:    "no email",
			content: "* Steve Durrheimer <s.durrheimer@gmail.com>\n",
			want: []Maintainer{
				{Name: "Steve Durrheimer", Email: "s.durrheimer@gmail.com"},
			},
		},
		{
			name:    "component prefix and multiple maintainers per line",
			content: "* `tsdb`: Ganesh Vernekar (<ganesh@grafana.com> / @codesome), Bartłomiej Płotka (<bwplotka@gmail.com> / @bwplotka)\n",
			want: []Maintainer{
				{Name: "Ganesh Vernekar", Email: "ganesh@grafana.com", Handle: "codesome"},
				{Name: "Bartłomiej Płotka", Email: "bwplotka@gmail.com", Handle: "bwplotka"},
			},
		},
		{
			name:    "email directly after name, handle in parens",
			content: "* Amy Super amy.super@grafana.com (@amy-super)\n",
			want: []Maintainer{
				{Name: "Amy Super", Email: "amy.super@grafana.com", Handle: "amy-super"},
			},
		},
		{
			name:    "handle without slash inside parens",
			content: "* `module`: Augustin Husson (<husson.augustin@gmail.com> @nexucis)\n",
			want: []Maintainer{
				{Name: "Augustin Husson", Email: "husson.augustin@gmail.com", Handle: "nexucis"},
			},
		},
		{
			name: "alumni section is skipped",
			content: "**Current Maintainers:**\n" +
				"- Ben Reedy (breed808) - breed808@breed808.com\n\n" +
				"**Alumni Contributors:**\n" +
				"- Brian Brazil (brian-brazil)\n" +
				"- Calle Pettersson (carlpett)\n",
			want: []Maintainer{
				{Name: "Ben Reedy", Email: "breed808@breed808.com", Handle: "breed808"},
			},
		},
		{
			name: "markdown profile links, plain-text alumni heading skipped",
			content: "Maintainers in alphabetical order\n\n" +
				"* [Ben Reedy](https://github.com/breed808) - breed808@breed808.com\n" +
				"* [Jan-Otto Kröpke](https://github.com/jkroepke) - github@jkroepke.de\n\n" +
				"Alumni\n\n" +
				"* [Brian Brazil](https://github.com/brian-brazil)\n",
			want: []Maintainer{
				{Name: "Ben Reedy", Email: "breed808@breed808.com", Handle: "breed808"},
				{Name: "Jan-Otto Kröpke", Email: "github@jkroepke.de", Handle: "jkroepke"},
			},
		},
		{
			name: "markdown link to another repo's file is not a maintainer",
			content: "to the maintainers specified in Alertmanager's\n" +
				"[MAINTAINERS.md](https://github.com/prometheus/alertmanager/blob/main/MAINTAINERS.md)\n" +
				"file for documentation about the Alertmanager.\n",
			want: nil,
		},
		{
			name: "markdown heading sections are not parsed as maintainers",
			content: "# Maintainers\n\n" +
				"* Julien Pivotto <roidelapluie@prometheus.io> @roidelapluie\n",
			want: []Maintainer{
				{Name: "Julien Pivotto", Email: "roidelapluie@prometheus.io", Handle: "roidelapluie"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseMaintainers(tt.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMaintainers() mismatch\n got: %#v\nwant: %#v", got, tt.want)
			}
		})
	}
}

func TestLoadEmailMapping(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "email-to-github.csv")
	content := "# a comment\n" +
		"email,github_id\n" +
		"superq@gmail.com,SuperQ\n" +
		"  Julius.Volz@gmail.com , juliusv \n" + // whitespace + mixed case
		"\n" +
		"# trailing comment\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := loadEmailMapping(path)
	if err != nil {
		t.Fatalf("loadEmailMapping: %v", err)
	}
	want := map[string]string{
		"superq@gmail.com":      "SuperQ",
		"julius.volz@gmail.com": "juliusv",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("loadEmailMapping() = %#v, want %#v", got, want)
	}

	// A missing file is not an error.
	got, err = loadEmailMapping(filepath.Join(dir, "does-not-exist.csv"))
	if err != nil {
		t.Fatalf("missing file should not error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("missing file should yield empty map, got %#v", got)
	}
}
