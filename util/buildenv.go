package util

import (
	"os"
)

func IsRunOnTravisCI() bool {
	_, ok := os.LookupEnv("TRAVIS")
	return ok
}

func IsRunOnGitLabCI() bool {
	_, ok := os.LookupEnv("GITLAB_CI")
	return ok
}
