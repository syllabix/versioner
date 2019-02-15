package semver

import (
	"fmt"
	"testing"
)

func Test_numtype_String(t *testing.T) {
	tests := []struct {
		name  string
		n     Type
		wantS string
	}{
		{
			name:  "Major",
			n:     Major,
			wantS: "Major",
		},
		{
			name:  "Minor",
			n:     Minor,
			wantS: "Minor",
		},
		{
			name:  "Patch",
			n:     Patch,
			wantS: "Patch",
		},
		{
			name:  "Pre Release",
			n:     PreRelease,
			wantS: "Pre Release",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := tt.n.String(); gotS != tt.wantS {
				t.Errorf("numtype.String() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type fields struct {
		major      int
		minor      int
		patch      int
		prerelease string
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields: fields{
				major: 3,
				minor: 12,
				patch: 123,
			},
			want: "3.12.123",
		},
		{
			fields: fields{
				major: 0,
				minor: 0,
				patch: 1,
			},
			want: "0.0.1",
		},
		{
			fields: fields{
				major:      1,
				minor:      0,
				patch:      1,
				prerelease: "rc.12",
			},
			want: "1.0.1-rc.12",
		},
	}
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("#%d", testnum), func(t *testing.T) {
			v := Version{
				major:      tt.fields.major,
				minor:      tt.fields.minor,
				patch:      tt.fields.patch,
				prerelease: tt.fields.prerelease,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Major(t *testing.T) {
	type fields struct {
		major      int
		minor      int
		patch      int
		prerelease string
	}
	tests := []struct {
		fields fields
		want   int
	}{
		{
			fields: fields{
				major: 3,
				minor: 12,
				patch: 123,
			},
			want: 3,
		},
		{
			fields: fields{
				major: 0,
				minor: 0,
				patch: 1,
			},
			want: 0,
		},
		{
			fields: fields{
				major:      12,
				minor:      0,
				patch:      1,
				prerelease: "rc.12",
			},
			want: 12,
		},
	}
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("#%d", testnum), func(t *testing.T) {
			v := Version{
				major:      tt.fields.major,
				minor:      tt.fields.minor,
				patch:      tt.fields.patch,
				prerelease: tt.fields.prerelease,
			}
			if got := v.Major(); got != tt.want {
				t.Errorf("Version.Major() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Minor(t *testing.T) {
	type fields struct {
		major      int
		minor      int
		patch      int
		prerelease string
	}
	tests := []struct {
		fields fields
		want   int
	}{
		{
			fields: fields{
				major: 3,
				minor: 12,
				patch: 123,
			},
			want: 12,
		},
		{
			fields: fields{
				major: 0,
				minor: 42,
				patch: 1,
			},
			want: 42,
		},
		{
			fields: fields{
				major:      12,
				minor:      0,
				patch:      1,
				prerelease: "rc.12",
			},
			want: 0,
		},
	}
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("#%d", testnum), func(t *testing.T) {
			v := Version{
				major:      tt.fields.major,
				minor:      tt.fields.minor,
				patch:      tt.fields.patch,
				prerelease: tt.fields.prerelease,
			}
			if got := v.Minor(); got != tt.want {
				t.Errorf("Version.Minor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Patch(t *testing.T) {
	type fields struct {
		major      int
		minor      int
		patch      int
		prerelease string
	}
	tests := []struct {
		fields fields
		want   int
	}{
		{
			fields: fields{
				major: 3,
				minor: 12,
				patch: 123,
			},
			want: 123,
		},
		{
			fields: fields{
				major: 0,
				minor: 42,
				patch: 1,
			},
			want: 1,
		},
		{
			fields: fields{
				major:      12,
				minor:      0,
				patch:      0,
				prerelease: "rc.12",
			},
			want: 0,
		},
	}
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("#%d", testnum), func(t *testing.T) {
			v := Version{
				major:      tt.fields.major,
				minor:      tt.fields.minor,
				patch:      tt.fields.patch,
				prerelease: tt.fields.prerelease,
			}
			if got := v.Patch(); got != tt.want {
				t.Errorf("Version.Patch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_PreRelease(t *testing.T) {
	type fields struct {
		major      int
		minor      int
		patch      int
		prerelease string
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields: fields{
				major: 3,
				minor: 12,
				patch: 123,
			},
			want: "",
		},
		{
			fields: fields{
				major:      0,
				minor:      42,
				patch:      1,
				prerelease: "alpha-1.0.1",
			},
			want: "alpha-1.0.1",
		},
		{
			fields: fields{
				major:      12,
				minor:      0,
				patch:      1,
				prerelease: "rc.12",
			},
			want: "rc.12",
		},
	}
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("#%d", testnum), func(t *testing.T) {
			v := Version{
				major:      tt.fields.major,
				minor:      tt.fields.minor,
				patch:      tt.fields.patch,
				prerelease: tt.fields.prerelease,
			}
			if got := v.PreRelease(); got != tt.want {
				t.Errorf("Version.PreRelease() = %v, want %v", got, tt.want)
			}
		})
	}
}
