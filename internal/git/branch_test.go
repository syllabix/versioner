package git

import (
	"io"
	"strings"
	"testing"
)

func Test_currentBranch(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should return current branch when it is first in output",
			args: args{
				r: strings.NewReader(outputA),
			},
			want: "feature/MYFEAT-1234",
		},
		{
			name: "Should return current branch when it is last in output",
			args: args{
				r: strings.NewReader(outputB),
			},
			want: "feature/MYFEAT-82943",
		},
		{
			name: "Should return current branch when it is output somewhere in the middle of the list",
			args: args{
				r: strings.NewReader(outputC),
			},
			want: "develop",
		},
		{
			name: "Should return an error if current branch is not flagged with a * ",
			args: args{
				r: strings.NewReader(outputD),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Should return current branch when list is long",
			args: args{
				r: strings.NewReader(outputF),
			},
			want:    "feature/MY-WORKING-BRANCH",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := currentBranch(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("currentBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("currentBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

//////////////
// TEST OUTPUT
//////////////

const (
	outputA = `
* feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321`

	outputB = `
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
* feature/MYFEAT-82943`

	outputC = `
feature/MYFEAT-1234
bugfix/dropdown-broken
* develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-82943`

	outputD = `
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-82943`

	outputE = ""

	outputF = `
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-82943
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-82943
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-82943
feature/MYFEAT-1234
bugfix/dropdown-broken
develop
master
feature/MYFEAT-54321
* feature/MY-WORKING-BRANCH
task/ATASK-9999
bugfix/it-is-not-working
feature/MYFEAT-82943
`
)
