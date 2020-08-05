package util

import (
	"os"
)

// IsRunOnTravisCI checks if running on TRAVIS CI environment.
func IsRunOnTravisCI() bool {
	_, ok := os.LookupEnv("TRAVIS")
	return ok
}

// IsRunOnGitLabCI checks if running on GitLab CI/CD environment.
func IsRunOnGitLabCI() bool {
	_, ok := os.LookupEnv("GITLAB_CI")
	return ok
}
