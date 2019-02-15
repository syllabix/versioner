package main

import (
	"flag"
	"os"

	"github.com/syllabix/versioner/internal/diagnostic"

	"github.com/fatih/color"
	"github.com/syllabix/versioner"
)

func main() {

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "print the current version")
	flag.Parse()

	if flag.Arg(0) == "version" {
		showVersion = true
	}

	if showVersion {
		color.Cyan("Versioner v%s", diagnostic.AppVersion)
		os.Exit(0)
	}

	version, err := versioner.Next()
	if err != nil {
		color.Red("unable to determine next version\n%v", err)
	}
	color.Green("%s", version)
}
