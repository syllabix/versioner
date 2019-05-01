package tag

import (
	"errors"
	"os/exec"
	"strings"
)

// GetLatest attempts to find and return the latest
// version of a project by retrieving a repositories
// latest annotated tag, returning an empty version and
// non nil error in the event of failure
func GetLatest() (string, error) {
	cmd := exec.Command("git", "describe", "--abbrev=0")
	tag, err := cmd.Output()
	if err != nil {
		return "", err
	}
	v := strings.TrimRight(string(tag), "\r\n")
	return v, nil
}

// GetVersionPriorTo will return the version prior to the provided version.
// If the provided version, or the prior version cannot be found, a non nil error
// will be returned
func GetVersionPriorTo(version string) (string, error) {
	cmd := exec.Command("git", "tag")
	tags, err := cmd.Output()
	if err != nil {
		return "", err
	}
	versions := strings.Split(string(tags), "\n")
	var prior string
	if len(versions) > 0 {
		for i, ver := range versions {
			if ver == version {
				if (i - 1) < 0 {
					return "", errors.New("no prior version found")
				}
				prior = versions[i-1]
				break
			}
		}
	}
	if prior == "" {
		return "", errors.New("the provided version was not found")
	}
	return prior, nil
}
