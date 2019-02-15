package git

import "os/exec"

// LatestCommitHash returns the latest commit hash when run within
// a working git repository. It will return an empty string and non nil
// error value on failure
func LatestCommitHash() (string, error) {
	cmd := exec.Command("git", "log", "--oneline", "--name-status HEAD^..HEAD")
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(b[0:8]), nil
}
