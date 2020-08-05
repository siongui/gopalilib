package util

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint pretty-prints Go variable (struct, map, array, slice, etc.).
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// LocalPrintln will not print on Travis CI or GitLab CI/CD environment.
func LocalPrintln(a ...interface{}) {
	if IsRunOnTravisCI() {
		return
	}
	if IsRunOnGitLabCI() {
		return
	}
	fmt.Println(a...)
}
