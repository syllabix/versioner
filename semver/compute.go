package semver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/syllabix/versioner/commit"
	"github.com/syllabix/versioner/internal/git"
)

var curhash func() (string, error)

// set up this way to simplify testing
func init() {
	curhash = git.LatestCommitHash
}

// ComputeNext will compute the next semantic version
func ComputeNext(v Version, commits []commit.Message) (Version, error) {

	maj := v.major
	min := v.minor
	ptch := v.patch

	var minorBumped bool
sumloop:
	for _, c := range commits {
		switch c.Type {
		case commit.Major:
			maj++
			min = 0
			ptch = 0
			break sumloop
		case commit.Minor:
			min = v.minor + 1
			ptch = 0
			minorBumped = true
		case commit.Patch:
			if !minorBumped {
				ptch++
			}
		}
	}

	return Version{
		major: maj,
		minor: min,
		patch: ptch,
	}, nil
}

// ComputeNextPreRelease attempts to provide a meaningful update
// to a prelease version, considering the following rules:
//
// if pre release version is semantically versioned itself - compute
// next version based on commit range
// if pre release version is a prefix terminated by a "." modify the suffix by:
//   - incrementing it by 1 if it is a number
//   - computing the next semantic version if it is a valid semantic version number
//   - replacing it with the current commit hash
// if none of the above can be determined - concatenate the current commit hash
func ComputeNextPreRelease(v Version, commits []commit.Message) (Version, error) {
	if len(v.prerelease) > 0 {
		pre, err := nextPreRelease(v.prerelease, commits)
		if err != nil {
			return Version{}, err
		}
		v.prerelease = pre
		return v, nil
	}

	return Version{}, fmt.Errorf("the provided version did not contain a pre release: %s", v)
}

func nextPreRelease(vstr string, commits []commit.Message) (string, error) {

	v, err := Parse(vstr)
	if err == nil {
		return nextSemanticPreRelease(v, commits)
	}

	tks := strings.SplitAfterN(vstr, ".", 2)
	if len(tks) == 2 {
		num, err := strconv.Atoi(tks[1])
		if err == nil {
			return fmt.Sprintf("%s%d", tks[0], num+1), nil
		}

		nxver, err := Parse(tks[1])
		if err == nil {
			nxStrVer, err := nextSemanticPreRelease(nxver, commits)
			return fmt.Sprintf("%s%s", tks[0], nxStrVer), err
		}

		hash, err := curhash()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s%s", tks[0], hash), nil
	}

	hash, err := curhash()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", vstr, hash), nil
}

func nextSemanticPreRelease(v Version, commits []commit.Message) (string, error) {
	v.prerelease = "" //hack - for now ensure no "pre-release of a pre-release" value is re computed avoiding infinite loop
	n, err := ComputeNext(v, commits)
	return n.String(), err
}
