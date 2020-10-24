package main

import (
	"errors"
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

	var (
		withLog     bool
		printscopes bool
		showVersion bool
		prerelease  string
		tagPrefix   string
		strictMode  bool
	)

	flag.BoolVar(&withLog, "with-changelog", false, "generates a change log and writes it to CHANGELOG.md")
	flag.BoolVar(&showVersion, "v", false, "print the current version of the binary")
	flag.BoolVar(&printscopes, "print-scopes", false, "print all found scopes at the provided version, falling back the current working version if not provided")
	flag.StringVar(&prerelease, "pre-release", "", "bumps the version as a new version for the provided pre-release")
	flag.StringVar(&tagPrefix, "tag-prefix", "", "match latest tag based on prefix filter")
	flag.BoolVar(&strictMode, "strict", false, "fail when there is no version from a tag or version can not be parsed")
	flag.Parse()

	if flag.Arg(0) == "version" {
		showVersion = true
	}

	if showVersion {
		color.Cyan("Versioner %s", diagnostic.AppVersion)
		os.Exit(0)
	}

	var latest string
	versionTag, err := tag.GetLatest(tagPrefix)
	if err == nil {
		latest = versionTag
	} else {
		if strictMode {
			fail(errors.New("couldn't compute last version tag"))
		} else {
			// no tags were found - assuming new project
			versionTag = "0.0.0"
		}
	}

	version, err := semver.Parse(versionTag)
	if err != nil {
		if strictMode {
			fail(err)
		} else {
			version, _ = semver.Parse("0.0.0")
		}
	}

	if printscopes {
		scopePrinter(latest)
		return
	}

	msgs, err := commit.MessagesInRange("HEAD", latest)
	if err != nil {
		fail(err)
	}

	var vnext semver.Version
	if len(prerelease) > 0 {
		preversion, err := semver.Parse(prerelease)
		if err != nil {
			fail(err)
		}

		if preversion.Major() != version.Major() ||
			preversion.Minor() != version.Minor() ||
			preversion.Patch() != version.Patch() {
			version = preversion
		}

		vnext, err = semver.ComputeNextPreRelease(version, msgs)
		if err != nil {
			fail(err)
		}
	} else {
		vnext, err = semver.ComputeNext(version, msgs)
		if err != nil {
			fail(err)
		}
	}

	if withLog {
		generator := changelog.NewGenerator(vnext, msgs)

		f, err := os.Create("CHANGELOG.md")
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
