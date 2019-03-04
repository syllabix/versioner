package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/syllabix/versioner/changelog"
	"github.com/syllabix/versioner/commit"
	"github.com/syllabix/versioner/semver"
	"github.com/syllabix/versioner/tag"

	"github.com/syllabix/versioner/internal/diagnostic"

	"github.com/fatih/color"
)

var (
	showVersion bool
	logtitle    string
	nolog       bool
	printscopes bool
)

func main() {

	flag.BoolVar(&showVersion, "v", false, "print the current version")
	flag.StringVar(&logtitle, "o", "CHANGELOG.md", "sets the name of the output file")
	flag.BoolVar(&nolog, "nolog", false, "disable generating change log")
	flag.BoolVar(&printscopes, "print-scopes", true, "print all found scopes delimited by space since most recent annotated tag")
	flag.Parse()

	if flag.Arg(0) == "version" {
		showVersion = true
	}

	if showVersion {
		color.Cyan("Versioner %s", diagnostic.AppVersion)
		os.Exit(0)
	}

	var latest string
	versionTag, err := tag.GetLatest()
	if err == nil {
		latest = versionTag
	} else {
		// no tags were found - assuming new project
		versionTag = "0.0.0"
	}

	version, err := semver.Parse(versionTag)
	if err != nil {
		color.Red("%+v", err)
		os.Exit(1)
	}

	msgs, err := commit.MessagesInRange("HEAD", latest)
	if err != nil {
		color.Red("%+v", err)
		os.Exit(1)
	}

	if printscopes {
		printScopes(msgs)
		return
	}

	vnext, err := semver.ComputeNext(version, msgs)
	if err != nil {
		color.Red("%+v", err)
		os.Exit(1)
	}

	if !nolog {
		generator := changelog.NewGenerator(vnext, msgs)

		f, err := os.Create(logtitle)
		if err != nil {
			color.Red("%+v", err)
			os.Exit(1)
		}
		defer f.Close()

		err = generator.Generate(f)
		if err != nil {
			color.Red("%+v", err)
			os.Exit(1)
		}
	}

	color.Green("%s", vnext)
}

func printScopes(msgs []commit.Message) {
	scopes := map[string]struct{}{}
	var builder strings.Builder
	for _, m := range msgs {
		if len(m.Scope) < 1 {
			continue
		}
		_, printed := scopes[m.Scope]
		if !printed {
			builder.WriteString(m.Scope + " ")
		}
		scopes[m.Scope] = struct{}{}
	}
	fmt.Println(builder.String())
}
