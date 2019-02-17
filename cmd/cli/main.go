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

func main() {

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "print the current version")

	var logtitle string
	flag.StringVar(&logtitle, "-o", "CHANGELOG.md", "sets the name of the output file")

	flag.Parse()

	if flag.Arg(0) == "version" {
		showVersion = true
	}

	if showVersion {
		color.Cyan("Versioner %s", diagnostic.AppVersion)
		os.Exit(0)
	}

	var to string
	vTag, err := tag.GetLatest()
	if err == nil {
		to = vTag
	} else {
		// no tags were found - assuming new project
		vTag = "0.0.0"
	}

	version, err := semver.Parse(vTag)
	if err != nil {
		color.Red("%+v", err)
		os.Exit(1)
	}

	msgs, err := commit.MessagesInRange("HEAD", to)
	if err != nil {
		color.Red("%+v", err)
		os.Exit(1)
	}

	vnext, err := semver.ComputeNext(version, msgs)
	if err != nil {
		color.Red("%+v", err)
		os.Exit(1)
	}

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

	color.Green("%s", vnext)
}
