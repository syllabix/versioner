package semver

import (
	"reflect"
	"testing"
	"time"

	"github.com/syllabix/versioner/commit"
)

func TestComputeNext(t *testing.T) {
	time.Sleep(time.Millisecond * 1)
	curhash = func() (string, error) {
		return "8d4c163", nil
	}

	for _, tt := range ctests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ComputeNext(tt.args.v, tt.args.commits)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeNext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkComputeNext(b *testing.B) {
	for _, test := range ctests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ComputeNext(test.args.v, test.args.commits)
			}
		})
	}
}

type cargs struct {
	v       Version
	commits []commit.Message
}

var ctests = []struct {
	name    string
	args    cargs
	want    Version
	wantErr bool
}{
	{
		name: "patch version",
		args: cargs{
			v: Version{
				major: 0,
				minor: 1,
				patch: 5,
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.NoImpact,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.NoImpact,
				},
				commit.Message{
					Type: commit.Patch,
				},
			},
		},
		want: Version{
			major: 0,
			minor: 1,
			patch: 8,
		},
		wantErr: false,
	},
	{
		name: "minor version",
		args: cargs{
			v: Version{
				major: 0,
				minor: 8,
				patch: 5,
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Minor,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Minor,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Minor,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
			},
		},
		want: Version{
			major: 0,
			minor: 9,
			patch: 0,
		},
		wantErr: false,
	},
	{
		name: "major version",
		args: cargs{
			v: Version{
				major: 2,
				minor: 8,
				patch: 5,
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Minor,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Minor,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Major,
				},
				commit.Message{
					Type: commit.Major,
				},
				commit.Message{
					Type: commit.Major,
				},
				commit.Message{
					Type: commit.Minor,
				},
			},
		},
		want: Version{
			major: 3,
			minor: 0,
			patch: 0,
		},
		wantErr: false,
	},
	{
		name: "semantic pre-release",
		args: cargs{
			v: Version{
				major:      2,
				minor:      8,
				patch:      5,
				prerelease: "0.0.1",
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Minor,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Minor,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Minor,
				},
			},
		},
		want: Version{
			major:      2,
			minor:      8,
			patch:      5,
			prerelease: "0.1.0",
		},
		wantErr: false,
	},
	{
		name: "prefixed semantic pre-release",
		args: cargs{
			v: Version{
				major:      2,
				minor:      8,
				patch:      5,
				prerelease: "rc.0.0.1",
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
			},
		},
		want: Version{
			major:      2,
			minor:      8,
			patch:      5,
			prerelease: "rc.0.0.5",
		},
		wantErr: false,
	},
	{
		name: "incremented prefixed pre-release",
		args: cargs{
			v: Version{
				major:      2,
				minor:      8,
				patch:      5,
				prerelease: "rc.12",
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
			},
		},
		want: Version{
			major:      2,
			minor:      8,
			patch:      5,
			prerelease: "rc.13",
		},
		wantErr: false,
	},
	{
		name: "prefixed pre-release fallback to hash",
		args: cargs{
			v: Version{
				major:      2,
				minor:      8,
				patch:      5,
				prerelease: "rc.asasf",
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
			},
		},
		want: Version{
			major:      2,
			minor:      8,
			patch:      5,
			prerelease: "rc.8d4c163",
		},
		wantErr: false,
	},
	{
		name: "unknown, concat hash",
		args: cargs{
			v: Version{
				major:      2,
				minor:      8,
				patch:      5,
				prerelease: "rc1234md9",
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
			},
		},
		want: Version{
			major:      2,
			minor:      8,
			patch:      5,
			prerelease: "rc1234md9.8d4c163",
		},
		wantErr: false,
	},
	{
		name: "unknown, concat hash",
		args: cargs{
			v: Version{
				major:      2,
				minor:      8,
				patch:      5,
				prerelease: "alpha_2_2_2019",
			},
			commits: []commit.Message{
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
				commit.Message{
					Type: commit.Patch,
				},
			},
		},
		want: Version{
			major:      2,
			minor:      8,
			patch:      5,
			prerelease: "alpha_2_2_2019.8d4c163",
		},
		wantErr: false,
	},
}
