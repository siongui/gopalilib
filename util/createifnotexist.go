package util

import (
	"os"
	"path"
)

// CreateDirIfNotExist creates the directory containing the file if the
// directory does not exist, given the path of the file. Similar to the shell
// command ``mkdir -p``. This method does not return error. If the directory
// cannot be made, this method panics.
func CreateDirIfNotExist(filepath string) {
	dir := path.Dir(filepath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
