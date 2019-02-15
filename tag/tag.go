package tag

import (
	"os/exec"
	"strings"
)

// GetLatest attempts to find and return the latest
// version of a project by retrieving a repositories
// latest annotated tag, returning an empty version and
// non nil error in the event of failure
func GetLatest() (string, error) {
	cmd := exec.Command("git", "describe")
	tag, err := cmd.Output()
	if err != nil {
		return "", err
	}
	v := strings.TrimRight(string(tag), "\r\n")
	return v, nil
}
