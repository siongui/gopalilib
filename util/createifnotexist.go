package util

import (
	"os"
	"path"
)

// Given the path of a file, create the directory containing the file if it does
// not exist. Similar to ``mkdir -p`` of shell command
func CreateDirIfNotExist(filepath string) {
	dir := path.Dir(filepath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
