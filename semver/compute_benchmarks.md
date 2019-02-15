### benchmarks for the semver parse function:

_let's do better_ :)

Sun Feb 17 2019

goos: darwin
goarch: amd64
pkg: github.com/syllabix/versioner/semver

BenchmarkComputeNext/patch_version-8    30000000	        50.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkComputeNext/minor_version-8    30000000	        54.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkComputeNext/major_version-8    30000000	        53.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkComputeNext/
semantic_pre-release-8  	            1000000	      1314 ns/op	    4153 B/op	       6 allocs/op
BenchmarkComputeNext/
prefixed_semantic_pre-release-8         500000	      3317 ns/op	    8499 B/op	      15 allocs/op
BenchmarkComputeNext/
incremented_prefixed_pre-release-8      1000000	      1480 ns/op	    4289 B/op	       8 allocs/op
BenchmarkComputeNext/
prefixed_pre-release_fallback_to_hash-8 200	   7439225 ns/op	   55365 B/op	     141 allocs/op
BenchmarkComputeNext/
unknown,_concat_hash-8                  200	   8332962 ns/op	   51070 B/op	     136 allocs/op
BenchmarkComputeNext/
unknown,_concat_hash#01-8               200	   8571329 ns/op	   51075 B/op	     136 allocs/op