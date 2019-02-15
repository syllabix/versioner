package git

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
	"strings"
)

// CurrentBranch will try to return the name of git branch
// currently checked out. In the event of failure, an empty
// string and non nil error value will be returned
func CurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch")
	out, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	defer cmd.Wait()

	return currentBranch(out)
}

func currentBranch(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	var branch string
scan:
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && string(line[0]) == "*" {
			branch = strings.Trim(line, "* \r\n")
			break scan
		}
	}
	if branch == "" {
		return "", errors.New("unable determine the currently checked out branch")
	}
	return branch, scanner.Err()
}
