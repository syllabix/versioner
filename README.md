# `versioner`
[![Go Report Card](https://goreportcard.com/badge/github.com/syllabix/versioner)](https://goreportcard.com/report/github.com/syllabix/versioner)

compute semantic versions and generate changelogs with 3 opinions:

1. your project is versioned using `git`.
2. commit messages in the repository follow the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0-beta.3/) standard
3. releases are tagged with [annotated tags](https://git-scm.com/book/en/v2/Git-Basics-Tagging)

## `versioner` obectives

a. be a simple tool to use

b. Follow two conventions

1. [Conventional Commits v1.0.0-beta.3](https://www.conventionalcommits.org/en/v1.0.0-beta.3/) - describes the syntactic convention to use in a commit message (which this program attempts to parse and derive a meaningful version number from)

2. [Semver 2.0.0](https://semver.org/) - a known convention for providing meaningful versions to software.

c. Derive an accurate semantic version number from a `git` managed history that uses tags to mark versions.

d. Output a meaningful changelog.


### usage

this project is in a pre release state, and can be installed in two ways:

1. install from source:
you will need the [Go tooling installed](https://golang.org/dl/). then run:

    `go get -u github.com/syllabix/versioner/cmd/versioner`

2. Download a binary from the from the latest built release [here](https://github.com/syllabix/versioner/releases)

Once downloaded - and ensuring the binary is in your system `PATH` - simply navigate into a git repostory and run `versioner`. if your repository is using conventional commits, a meaningful version should be output

```
  // possible flags
  -pre-release string
        bumps the version as a new version for the provided pre-release
  -print-scopes
        print all found scopes at the provided version, falling back the current working version if not provided
  -v    print the current version of the binary
  -with-changelog
        generates a change log and writes it to CHANGELOG.md
```

#### pre releases
pre releases are handled with versioner by explicitly passing in the upcoming version with it's respective zero value pre release template. versioner will then try to reasonably bump the pre release version with conventional commits as long as the major.minor.patch passed to the `pre-release` flag is the same as the previous annotated tag in your git history

for example:
```
// upcoming release is 0.2.0 - we would like bump on a release candidate with it's own semantic version

for example: 0.2.0-rc.0.0.1

// then work starts for 0.2.0
git commit -m "task: awesome"
git commit -m "task: great"

// time to release pre release
versioner -pre-release 0.2.0-rc.0.0.0

// outputs =>
0.2.0-rc.0.0.2

// more work happens
git commit -m "feat: super great"

// time to release
versioner -pre-release 0.2.0-rc.0.0.0

// outputs =>
0.2.0-rc.0.1.0
```

### roadmap
1. installable via homebrew
2. installable via chocolatey
3. installable via apt-get