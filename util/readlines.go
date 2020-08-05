// Package util provides utility func.
package util

import (
	"bufio"
	"os"
)

// ReadlinesFromFile reads lines from the file.
func ReadlinesFromFile(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}
