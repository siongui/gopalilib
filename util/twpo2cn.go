package util

import (
	"os"
	"path/filepath"
	"strings"
)

// TwPoToCN creates the zh_CN PO file from the given zh_TW PO file.
func TwPoToCN(twPOPath, cnPOPath string) {
	os.MkdirAll(filepath.Dir(cnPOPath), 0755)
	fo, err := os.Create(cnPOPath)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	lines, err := ReadlinesFromFile(twPOPath)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "msgstr") {
			fo.Write([]byte(T2S(line)))
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
