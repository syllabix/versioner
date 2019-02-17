//go:generate statik -src=./template

package changelog

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"

	"github.com/rakyll/statik/fs"
	"github.com/syllabix/versioner/commit"
	"github.com/syllabix/versioner/semver"

	// This empty import is required to initialize the tempate data
	_ "github.com/syllabix/versioner/changelog/statik"
)

type Generator struct {
	v       semver.Version
	commits []commit.Message
}

// Generate a changelog writing log the provided writer
func (g Generator) Generate(w io.Writer) error {
	statikfs, err := fs.New()
	if err != nil {
		return err
	}

	f, err := statikfs.Open("/standard.md")
	if err != nil {
		return err
	}

	var buffer strings.Builder
	io.Copy(&buffer, f)

	tmpl, err := template.New("changelog").Parse(buffer.String())
	if err != nil {
		return err
	}

	d := Data{
		Version: g.v.String(),
		Date:    time.Now().Format("Jan 02 2006"),
	}

	contributors := make(map[string]bool)

	for _, c := range g.commits {
		contributors[c.Author] = true
		line := fmt.Sprintf("%s\n%s\n%s", c.Subject, c.Body, c.Footer)
		line = strings.Trim(line, " \n")
		switch c.Type {
		case commit.Major:
			d.BreakingChanges = append(d.BreakingChanges, line)
		case commit.Minor:
			d.Features = append(d.Features, line)
		case commit.Patch:
			d.Fixes = append(d.Fixes, line)
		}
	}

	for author := range contributors {
		d.Contributors = append(d.Contributors, author)
	}

	return tmpl.Execute(w, d)
}

// NewGenerator returns a Generator
func NewGenerator(v semver.Version, commits []commit.Message) Generator {
	return Generator{
		v:       v,
		commits: commits,
	}
}
