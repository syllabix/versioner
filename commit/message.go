package commit

import (
	"time"
)

// A Type represents what type of version update
// the commit impacts
type Type uint8

// Commit types
const (
	NoImpact Type = iota + 1
	Patch
	Minor
	Major
)

// Common conventional commit prefixes
// A prefix is the first portion of a conventional commit message or description
// body that identifies its type - and is positioned before
// an optional scope and ':'
//
// feat: my awesome feature
// ^--^
//  | -> "feat" is the prefix
const (
	Feat           = "feat"
	Fix            = "fix"
	Chore          = "chore"
	Improvement    = "improvement"
	BreakingChange = "BREAKING CHANGE"
)

// Message contains the type and content of a commit
type Message struct {
	Author  string
	Date    time.Time
	Type    Type
	Scope   string
	Subject string
	Body    string
	Footer  string
}
