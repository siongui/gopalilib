// Package i18n provides tools for Internationalization.
package i18n

import (
	"bufio"
	"fmt"
	"github.com/siongui/gojianfan"
	"os"
	"path/filepath"
	"strings"
)

func File2lines(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}

	return lines
}

// Create the zh_CN PO file from the given zh_TW PO file.
func TwPoToCN(twPOPath, cnPOPath string) {
	os.MkdirAll(filepath.Dir(cnPOPath), 0755)
	fo, err := os.Create(cnPOPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}
	defer fo.Close()

	for _, line := range File2lines(twPOPath) {
		if strings.HasPrefix(line, "msgstr") {
			fo.Write([]byte(gojianfan.T2S(line)))
		} else {
			if strings.Contains(line, "zh_TW") {
				fo.Write([]byte(strings.Replace(line, "zh_TW", "zh_CN", 1)))
			} else {
				fo.Write([]byte(line))
			}
		}
		fo.Write([]byte("\n"))
	}
}
