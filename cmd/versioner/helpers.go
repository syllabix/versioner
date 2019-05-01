package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/syllabix/versioner/commit"
	"github.com/syllabix/versioner/tag"
)

func fail(err error) {
	color.Red("%+v", err)
	os.Exit(1)
}

func scopePrinter(latest string) {
	args := flag.Args()
	v := "HEAD"
	n := len(args)
	if n > 0 {
		v = args[n-1]
	}

	if v == "HEAD" {
		msgs, err := commit.MessagesInRange("HEAD", latest)
		if err != nil {
			fail(err)
		}
		printScopes(msgs)
	} else {
		msgs := commitsForVersion(v)
		printScopes(msgs)
	}
}

func commitsForVersion(v string) []commit.Message {
	prior, err := tag.GetVersionPriorTo(v)
	if err != nil {
		color.Red("%+v", err)
		os.Exit(1)
	}
	msgs, err := commit.MessagesInRange(v, prior)
	if err != nil {
		color.Red("%+v", err)
		os.Exit(1)
	}
	return msgs
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
