package diagnostic

var (
	// AppVersion represents the current version of the binary
	AppVersion string
	// BuildTimestamp represents the build time of the binary
	BuildTimestamp string
	// CommitHash is the git commit hash the binary was built from
	CommitHash string
	// GoVersion - the version of go this binary was written in and built with
	GoVersion string
)
