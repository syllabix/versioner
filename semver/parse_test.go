package semver

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	for testnum, tt := range tests {
		t.Run(fmt.Sprintf("#%d", testnum), func(t *testing.T) {
			got, err := Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for _, test := range tests {
		b.Run(test.args.s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Parse(test.args.s)
			}
		})
	}
}

type args struct {
	s string
}

var tests = []struct {
	args    args
	want    Version
	wantErr bool
}{
	0: {
		args: args{
			s: "v0.12.345",
		},
		want: Version{
			patch: 345,
			minor: 12,
			major: 0,
		},
	},
	1: {
		args: args{
			s: "0.1.14",
		},
		want: Version{
			patch: 14,
			minor: 1,
			major: 0,
		},
	},
	2: {
		args: args{
			s: "v002.23.8",
		},
		want:    Version{},
		wantErr: true,
	},
	3: {
		args: args{
			s: "34t.asd.14",
		},
		want:    Version{},
		wantErr: true,
	},
	4: {
		args: args{
			s: "14.2",
		},
		want:    Version{},
		wantErr: true,
	},
	5: {
		args: args{
			s: "0.1.23-rc-1",
		},
		want: Version{
			patch:      23,
			minor:      1,
			major:      0,
			prerelease: "rc-1",
		},
	},
	6: {
		args: args{
			s: "0.1.23-beta.0.2.23",
		},
		want: Version{
			patch:      23,
			minor:      1,
			major:      0,
			prerelease: "beta.0.2.23",
		},
	},
	7: {
		args: args{
			s: "1.02.34234",
		},
		want:    Version{},
		wantErr: true,
	},
	8: {
		args: args{
			s: "1.2.03-rc-2",
		},
		want:    Version{},
		wantErr: true,
	},
	9: {
		args: args{
			s: "v0.1.23-beta.0.2.23",
		},
		want: Version{
			patch:      23,
			minor:      1,
			major:      0,
			prerelease: "beta.0.2.23",
		},
	},
	10: {
		args: args{
			s: "123.12.3423431232",
		},
		want: Version{
			major: 123,
			minor: 12,
			patch: 3423431232,
		},
	},
	11: {
		args: args{
			s: "adasgsdfsdf",
		},
		want:    Version{},
		wantErr: true,
	},
	12: {
		args: args{
			s: "4.4a.34b",
		},
		want:    Version{},
		wantErr: true,
	},
	13: {
		args: args{
			s: "0.1.0",
		},
		want: Version{
			major: 0,
			minor: 1,
			patch: 0,
		},
		wantErr: false,
	},
	14: {
		args: args{
			s: "12.6.4",
		},
		want: Version{
			major: 12,
			minor: 6,
			patch: 4,
		},
		wantErr: false,
	},
	15: {
		args: args{
			s: "12.6.4-beta.0.2.3",
		},
		want: Version{
			major:      12,
			minor:      6,
			patch:      4,
			prerelease: "beta.0.2.3",
		},
		wantErr: false,
	},
	16: {
		args: args{
			s: "2.3.004-beta.0.2.3",
		},
		want:    Version{},
		wantErr: true,
	},
	17: {
		args: args{
			s: "124152.323423.2342534534-rc.99923",
		},
		want: Version{
			major:      124152,
			minor:      323423,
			patch:      2342534534,
			prerelease: "rc.99923",
		},
		wantErr: false,
	},
	18: {
		args: args{
			s: "",
		},
		want:    Version{},
		wantErr: true,
	},
	19: {
		args: args{
			s: " ",
		},
		want:    Version{},
		wantErr: true,
	},
	20: {
		args: args{
			s: "\n\t",
		},
		want:    Version{},
		wantErr: true,
	},
	21: {
		args: args{
			s: "0.2.b",
		},
		want:    Version{},
		wantErr: true,
	},
	22: {
		args: args{
			s: "0.1c.b",
		},
		want:    Version{},
		wantErr: true,
	},
	23: {
		args: args{
			s: "z.zz.zzz",
		},
		want:    Version{},
		wantErr: true,
	},
}
