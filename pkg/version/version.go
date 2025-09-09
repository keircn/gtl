package version

import "fmt"

var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
	GoVersion = "unknown"
)

func BuildVersion() string {
	return fmt.Sprintf("gtl version %s\nCommit: %s\nBuild Date: %s\nGo Version: %s", Version, GitCommit, BuildDate, GoVersion)
}

func ShortVersion() string {
	return Version
}
