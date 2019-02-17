# `versioner`
[![Go Report Card](https://goreportcard.com/badge/github.com/syllabix/versioner)](https://goreportcard.com/report/github.com/syllabix/versioner)

compute semantic versions and generate changelogs

## motivation

after working on setting up automated change releases leveraging conventional commits, it became apparent that the best tools available are opinionated towards their respective langauges and/or runtimes.

this project is focused on building a semver calculator and changelog generator with only two opinions:

1. your project is versioned using `git`.
2. commit messages in the repository follow the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0-beta.3/) standard

## `versioner` obectives

a. be a simple tool to use

b. Follow two conventions

1. [Conventional Commits v1.0.0-beta.3](https://www.conventionalcommits.org/en/v1.0.0-beta.3/) - describes the syntactic convention to use in a commit message (which this program attempts to parse and derive a meaningful version number from)

2. [Semver 2.0.0](https://semver.org/) - a known convention for providing meaningful versions to software.

c. Derive an accurate semantic version number from a `git` managed history that uses tags to mark versions.

d. Output a meaningful changelog.


### usage

this project is in a pre release state, and is currently only available via installing from source, thus as a prerequisite you will need the [Go tooling installed](https://golang.org/dl/).

`go get -u github.com/syllabix/versioner`

Once installed - and ensuring the binary is in your system `PATH` - simply navigate into a git repostory and run `versioner`. if your repository is using conventional commits, a meaningful version should be output

### roadmap
0. Finish changelog generator
1. improve performance
2. handle pre release versions in a way more reflective of reasonable use cases semver standards.
3. installable via homebrew
4. installable via chocolatey
4. installable via apt-get