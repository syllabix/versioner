package semver

import (
	"fmt"
)

type Type int8

// Version number types
const (
	Major Type = iota + 1
	Minor
	Patch
	PreRelease
)

func (n Type) String() (s string) {
	switch n {
	case Major:
		s = "Major"
	case Minor:
		s = "Minor"
	case Patch:
		s = "Patch"
	case PreRelease:
		s = "Pre Release"
	default:
		s = "Invalid"
	}
	return
}

// A Version is a string representation of a SemVer number.
type Version struct {
	major      int
	minor      int
	patch      int
	prerelease string
}

func (v Version) String() string {
	vstr := fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
	if len(v.prerelease) > 0 {
		return fmt.Sprintf("%s-%s", vstr, v.prerelease)
	}
	return vstr
}

// Major portion of the version number
func (v Version) Major() int {
	return v.major
}

// Minor portion of the version number
func (v Version) Minor() int {
	return v.minor
}

// Patch portion of the version number
func (v Version) Patch() int {
	return v.patch
}

// PreRelease portion of the version number
func (v Version) PreRelease() string {
	return v.prerelease
}
