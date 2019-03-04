package commit

import (
	"bufio"
	"io"
	"strings"
	"time"
)

func parseMessages(r io.Reader) ([]Message, error) {
	scanner := bufio.NewScanner(r)

	var messages []Message

	var msg Message
	for scanner.Scan() {
		line := scanner.Text()

		if isEmpty(line) {
			continue
		}

		if isNewCommit(line) {
			messages = append(messages, msg)
			msg = Message{}
			continue
		}

		author, ok := getAuthor(line)
		if ok {
			msg.Author = author
			continue
		}

		date, ok := getDate(line)
		if ok {
			msg.Date = date
			continue
		}

		t, s, ok := getTypeAndSubject(line)
		if ok {
			msg.Type = t
			msg.Scope = s.scope
			if isEmpty(msg.Subject) {
				msg.Subject = s.subject
			} else if isEmpty(msg.Body) {
				msg.Body = s.subject
			} else {
				msg.Footer += s.subject
			}
			continue
		}

		if !isEmpty(line) {
			if isEmpty(msg.Body) {
				msg.Body = strings.Trim(line, "\r\t\n ")
				continue
			}
			msg.Footer += strings.Trim(line, "\r\t\n ")
		}
	}
	messages = append(messages, msg)
	return messages[1:], nil
}

func isNewCommit(line string) bool {
	line = strings.TrimSpace(line)
	if len(line) >= 6 {
		return line[0:6] == "commit"
	}
	return false
}

func getAuthor(line string) (string, bool) {
	info := detailify(line)
	if info.IsOfKind("author") {
		return info.Value(), true
	}
	return "", false
}

func getDate(line string) (time.Time, bool) {
	info := detailify(line)
	if info.IsOfKind("date") {
		t, err := time.Parse("Mon Jan 02 15:04:05 2006 -0700", info.Value())
		if err != nil {
			return time.Time{}, false
		}
		return t, true
	}
	return time.Time{}, false
}

type subscope struct {
	subject string
	scope   string
}

func getTypeAndSubject(line string) (Type, subscope, bool) {
	tokens := tokenize(line)
	var result subscope
	if len(tokens) == 2 {
		tk := strings.TrimRight(tokens[0], ":")
		parts := strings.Split(tk, "(")
		tk = parts[0]
		if len(parts) == 2 {
			result.scope = strings.TrimRight(parts[1], ")")
		}
		var ty Type
		switch tk {
		case Feat:
			ty = Minor
		case Fix:
			ty = Patch
		case BreakingChange:
			ty = Major
		default:
			// This is debatable - determine best way to handle these kind of commits
			ty = Patch
		}
		result.subject = strings.Trim(line, "\t\n\r ")
		return ty, result, true
	}
	return NoImpact, result, false
}

type details []string

func (d details) IsOfKind(k string) bool {
	return len(d) == 2 && strings.EqualFold(d[0], k+":")
}

func (d details) Value() string {
	if len(d) == 2 {
		return strings.Trim(d[1], "\t\n\r ")
	}
	return ""
}

func detailify(line string) details {
	return tokenize(line)
}

func tokenize(line string) []string {
	line = strings.TrimSpace(line)
	return strings.SplitAfterN(line, ":", 2)
}

func isEmpty(s string) bool {
	s = strings.TrimSpace(s)
	return s == ""
}
