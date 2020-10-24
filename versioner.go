package versioner

import (
	"github.com/syllabix/versioner/commit"
	"github.com/syllabix/versioner/semver"
	"github.com/syllabix/versioner/tag"
)

// Next will return the next semantic version derived
// from a git repository.
// A zero value version and non nil error will be returned in the event of failure
func Next(tagPrefix string) (semver.Version, error) {

	var to string
	vTag, err := tag.GetLatest(tagPrefix)
	if err == nil {
		to = vTag
	} else {
		// no tags were found - assuming new project
		vTag = "0.0.0"
	}

	version, err := semver.Parse(vTag)
	if err != nil {
		return semver.Version{}, err
	}

	msgs, err := commit.MessagesInRange("HEAD", to)
	if err != nil {
		return semver.Version{}, err
	}

	return semver.ComputeNext(version, msgs)
}
