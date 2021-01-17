// Package util provides helper methods for development.
// This package is for offline processing and not used in frontend code.
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

// IsRunOnGitHubActions checks if running on GitHub Actions environment.
func IsRunOnGitHubActions() bool {
	_, ok := os.LookupEnv("GITHUB_ACTIONS")
	return ok
}
