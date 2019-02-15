package commit

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_isNewCommit(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		args args
		want bool
	}{
		1: {
			args: args{
				line: "commit d1488da4e54f72326044956aef3fca28db026867 (HEAD -> feature/compute-semver)",
			},
			want: true,
		},
		2: {
			args: args{
				line: "Author: Dave Developer <d.developer@code.com>",
			},
			want: false,
		},
		3: {
			args: args{
				line: "Date:   Sat Feb 16 14:40:02 2019 +0100",
			},
			want: false,
		},
		4: {
			args: args{
				line: "     feat: adding commit package",
			},
			want: false,
		},
	}
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("TEST #%d", testnum), func(t *testing.T) {
			if got := isNewCommit(tt.args.line); got != tt.want {
				t.Errorf("isNewCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAuthor(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		args  args
		want  string
		want1 bool
	}{
		0: {
			args: args{
				line: "commit d1488da4e54f72326044956aef3fca28db026867 (HEAD -> feature/compute-semver)",
			},
			want:  "",
			want1: false,
		},
		1: {
			args: args{
				line: "Author: Cool Guy <cool.guy@gmail.com>",
			},
			want:  "Cool Guy <cool.guy@gmail.com>",
			want1: true,
		},
		2: {
			args: args{
				line: "aUthOr: Coder Girl <coder.girl@code.com>",
			},
			want:  "Coder Girl <coder.girl@code.com>",
			want1: true,
		},
		3: {
			args: args{
				line: "Date:   Sat Feb 16 14:40:02 2019 +0100",
			},
			want:  "",
			want1: false,
		},
		4: {
			args: args{
				line: "BREAKING CHANGE: just got serious",
			},
			want:  "",
			want1: false,
		},
		5: {
			args: args{
				line: "",
			},
			want:  "",
			want1: false,
		},
		6: {
			args: args{
				line: "aUthOr: mix:master:plus <mix@masterplus.com>",
			},
			want:  "mix:master:plus <mix@masterplus.com>",
			want1: true,
		},
	}
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("TEST #%d", testnum), func(t *testing.T) {
			got, got1 := getAuthor(tt.args.line)
			if got != tt.want {
				t.Errorf("isAuthor() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isAuthor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getDate(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		args  args
		want  time.Time
		want1 bool
	}{
		0: {
			args: args{
				line: "commit d1488da4e54f72326044956aef3fca28db026867 (HEAD -> feature/compute-semver)",
			},
			want:  time.Time{},
			want1: false,
		},
		1: {
			args: args{
				line: "Author: Tom Stoepker <tom.stoepker@gmail.com>",
			},
			want:  time.Time{},
			want1: false,
		},
		2: {
			args: args{
				line: "dAtE: Sat Mar 27 14:40:02 2019 +0100",
			},
			want: func() time.Time {
				t, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Sat Mar 27 14:40:02 2019 +0100")
				return t
			}(),
			want1: true,
		},
		3: {
			args: args{
				line: "Date:   Sat Feb 16 14:40:02 2019 +0100",
			},
			want: func() time.Time {
				t, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Sat Feb 16 14:40:02 2019 +0100")
				return t
			}(),
			want1: true,
		},
		4: {
			args: args{
				line: "BREAKING CHANGE: just got serious",
			},
			want:  time.Time{},
			want1: false,
		},
		5: {
			args: args{
				line: "",
			},
			want:  time.Time{},
			want1: false,
		},
		6: {
			args: args{
				line: "Date: Sat Feb 16 14:40:02 2019 +0100",
			},
			want: func() time.Time {
				t, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", "Sat Feb 16 14:40:02 2019 +0100")
				return t
			}(),
			want1: true,
		},
	}
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("TEST #%d", testnum), func(t *testing.T) {
			got, got1 := getDate(tt.args.line)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getDate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseMessages(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []Message
		wantErr bool
	}{
		{
			name: "Returns expected slice of messages",
			args: args{
				r: strings.NewReader(outputA),
			},
			want: []Message{
				Message{
					Author:  "Elmo MacAbee <elmo.macabee@tester.fun>",
					Date:    timeparse("Sat Feb 16 14:40:02 2019 +0100"),
					Type:    Major,
					Subject: "feat: adding commit package",
					Body:    "BREAKING CHANGE: just got serious",
				},
				Message{
					Author:  "Sandra Wahooey <s.wahoo@coder.co>",
					Date:    timeparse("Sat Feb 16 13:35:14 2019 +0100"),
					Type:    Minor,
					Subject: "feat: adding messages package",
				},
				Message{
					Author:  "Sandra Wahooey <s.wahoo@coder.co>",
					Date:    timeparse("Sat Feb 16 13:25:26 2019 +0100"),
					Type:    Patch,
					Subject: "chore: improve test utils",
				},
				Message{
					Author:  "Elmo MacAbee <elmo.macabee@tester.fun>",
					Date:    timeparse("Sat Feb 16 13:10:51 2019 +0100"),
					Type:    Patch,
					Subject: "docs: provide comment level clarity to test case",
				},
				Message{
					Author:  "Mike Mikerson <mike.mikerson@domainless.wha>",
					Date:    timeparse("Sat Feb 16 13:08:20 2019 +0100"),
					Type:    Patch,
					Subject: "chore: improve GetLatest func coverage",
					Body:    "fixes task #12345",
				},
				Message{
					Author:  "Elmo MacAbee <elmo.macabee@tester.fun>",
					Date:    timeparse("Sat Feb 16 13:01:26 2019 +0100"),
					Type:    Minor,
					Subject: "feat(coolthing): made this awesome thing so thingified",
					Body:    "closes task #912312",
				},
				Message{
					Author:  "Emma Zone <emma@zone.com>",
					Date:    timeparse("Sat Feb 16 12:58:53 2019 +0100"),
					Type:    Major,
					Subject: "feat: improving test for fetching latest version",
					Body:    "BREAKING CHANGE: yeah let's bump this one up majorly",
				},
				Message{
					Author:  "Emma Zone <emma@zone.com>",
					Date:    timeparse("Sat Feb 16 12:58:53 2019 +0100"),
					Type:    Major,
					Subject: "feat: improving test for fetching latest version",
					Body:    "yeah let's bump this one up majorly",
					Footer:  "BREAKING CHANGE: closes issue 2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMessages(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, msg := range tt.want {
				if !reflect.DeepEqual(got[i], msg) {
					t.Errorf("parseMessages() at index %d got =\n%+v\nwant =\n%+v", i, got[i], msg)
				}
			}
		})
	}
}

//////////////
// TEST OUTPUT FOR Test_parseMessages
//////////////

func timeparse(s string) time.Time {
	t, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", s)
	return t
}

const (
	outputA = `
	commit d1488da4e54f72326044956aef3fca28db026867 (HEAD -> feature/compute-semver)
	Author: Elmo MacAbee <elmo.macabee@tester.fun>
	Date:   Sat Feb 16 14:40:02 2019 +0100

		feat: adding commit package

		BREAKING CHANGE: just got serious

	commit f494fd95ce7001e88c10885bc6b175661d900933
	Author: Sandra Wahooey <s.wahoo@coder.co>
	Date:   Sat Feb 16 13:35:14 2019 +0100

		feat: adding messages package

	commit 8fa6cd1b2d9dffb6d6a68b7f5e413da8650e76d5 (tag: v0.1.0)
	Author: Sandra Wahooey <s.wahoo@coder.co>
	Date:   Sat Feb 16 13:25:26 2019 +0100

		chore: improve test utils

	commit c36074b0cd79a85cd5e70ebf524e243c1ee0a8d3
	Author: Elmo MacAbee <elmo.macabee@tester.fun>
	Date:   Sat Feb 16 13:10:51 2019 +0100

		docs: provide comment level clarity to test case

	commit 8970a24ab6b8939d8ac3b53c4a01d0eab3f3574a
	Author: Mike Mikerson <mike.mikerson@domainless.wha>
	Date:   Sat Feb 16 13:08:20 2019 +0100

		chore: improve GetLatest func coverage

		fixes task #12345

	commit 5e7f9fafda61da8b0e6f59cf8dc4d5a9abbb8fe6
	Author: Elmo MacAbee <elmo.macabee@tester.fun>
	Date:   Sat Feb 16 13:01:26 2019 +0100

		feat(coolthing): made this awesome thing so thingified

		closes task #912312

	commit e715b7b335291dd402a4dfb858a4f733b49a80dd
	Author: Emma Zone <emma@zone.com>
	Date:   Sat Feb 16 12:58:53 2019 +0100

		feat: improving test for fetching latest version

		BREAKING CHANGE: yeah let's bump this one up majorly

		commit e715b7b335291dd402a4dfb858a4f733b49a80dd
	Author: Emma Zone <emma@zone.com>
	Date:   Sat Feb 16 12:58:53 2019 +0100

		feat: improving test for fetching latest version

		yeah let's bump this one up majorly

		BREAKING CHANGE: closes issue 2
`
)
