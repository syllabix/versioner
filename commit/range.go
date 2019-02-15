package commit

import (
	"fmt"
	"os/exec"
)

// MessagesInRange will return commit messages between a valid start and end
// range identifier, returning an empty slice of Message and a non nil error
// on failure
func MessagesInRange(start, end string) ([]Message, error) {
	if isEmpty(start) {
		start = "HEAD"
	}

	var logrange string
	if isEmpty(end) {
		logrange = start
	} else {
		logrange = fmt.Sprintf("%s...%s", start, end)
	}

	cmd := exec.Command("git", "log", logrange)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return []Message{}, nil
	}

	err = cmd.Start()
	if err != nil {
		return []Message{}, nil
	}
	defer cmd.Wait()
	return parseMessages(out)
}
