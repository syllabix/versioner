package main

import (
	"flag"
	"os"

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

	flag.BoolVar(&showVersion, "v", false, "print the current version of the binary")
	flag.StringVar(&logtitle, "o", "CHANGELOG.md", "sets the name of the output file")
	flag.BoolVar(&nolog, "nolog", false, "disable generating change log")
	flag.BoolVar(&printscopes, "print-scopes", false, "print all found scopes at the provided version, falling back the current working version if not provided")
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
		fail(err)
	}

	if printscopes {
		scopePrinter(latest)
		return
	}

	msgs, err := commit.MessagesInRange("HEAD", latest)
	if err != nil {
		fail(err)
	}

	vnext, err := semver.ComputeNext(version, msgs)
	if err != nil {
		fail(err)
	}

	if !nolog {
		generator := changelog.NewGenerator(vnext, msgs)

		f, err := os.Create(logtitle)
		if err != nil {
			fail(err)
		}
		defer f.Close()

		err = generator.Generate(f)
		if err != nil {
			fail(err)
		}
	}

	color.Green("%s", vnext)
}
