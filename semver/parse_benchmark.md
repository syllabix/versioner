### benchmarks for the semver parse function:

_let's do better_ :)
```
Sun Feb 17 2019

goos: darwin
goarch: amd64
pkg: github.com/syllabix/versioner/semver

BenchmarkParse/v0.12.345-8         	 2000000	       951 ns/op	    4160 B/op	       7 allocs/op
BenchmarkParse/0.1.14-8            	 2000000	       926 ns/op	    4152 B/op	       6 allocs/op
BenchmarkParse/v002.23.8-8         	 1000000	      1043 ns/op	    4193 B/op	       4 allocs/op
BenchmarkParse/34t.asd.14-8        	 1000000	      1191 ns/op	    4233 B/op	       5 allocs/op
BenchmarkParse/14.2-8              	 2000000	       795 ns/op	    4144 B/op	       4 allocs/op
BenchmarkParse/0.1.23-rc-1-8       	 2000000	       922 ns/op	    4160 B/op	       8 allocs/op
BenchmarkParse/0.1.23-beta-0.2.23-8  2000000	       948 ns/op	    4176 B/op	       8 allocs/op
BenchmarkParse/1.02.34234-8          1000000	      1078 ns/op	    4233 B/op	       5 allocs/op
BenchmarkParse/1.2.03-rc-2-8         1000000	      1169 ns/op	    4241 B/op	       6 allocs/op
BenchmarkParse/v0.1.23-beta-0.2.23-8 2000000	       931 ns/op	    4176 B/op	       8 allocs/op
BenchmarkParse/123.12.3423431232-8   1000000	      1103 ns/op	    4200 B/op	       9 allocs/op
BenchmarkParse/adasgsdfsdf-8         1000000	      1028 ns/op	    4225 B/op	       4 allocs/op
BenchmarkParse/4.4a.34b-8            1000000	      1112 ns/op	    4241 B/op	       6 allocs/op
BenchmarkParse/0.1.0-8               2000000	       914 ns/op	    4144 B/op	       4 allocs/op
BenchmarkParse/12.6.4-8              2000000	       909 ns/op	    4152 B/op	       5 allocs/op
BenchmarkParse/12.6.4-beta-0.2.3-8   2000000	       940 ns/op	    4192 B/op	       8 allocs/op
BenchmarkParse/2.3.004-beta-0.2.3-8  1000000	      1186 ns/op	    4209 B/op	       6 allocs/op
BenchmarkParse/                      1000000	      1416 ns/op	    4216 B/op	      11 allocs/op
124152.323423.2342534534-rc-99923-8
BenchmarkParse/#00-8                 2000000	       784 ns/op	    4128 B/op	       2 allocs/op
BenchmarkParse/_-8                   2000000	       777 ns/op	    4176 B/op	       3 allocs/op
BenchmarkParse/__-8                  1000000	      1101 ns/op	    4225 B/op	       4 allocs/op
BenchmarkParse/0.2.b-8               1000000	      1050 ns/op	    4192 B/op	       5 allocs/op
BenchmarkParse/0.1c.b-8              1000000	      1305 ns/op	    4241 B/op	       6 allocs/op
BenchmarkParse/z.zz.zzz-8            1000000	      1175 ns/op	    4225 B/op	       4 allocs/op
```