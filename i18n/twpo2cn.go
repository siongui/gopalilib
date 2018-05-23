// Package i18n provides tools for Internationalization.
package i18n

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/siongui/gopalilib/util"
)

// Create the zh_CN PO file from the given zh_TW PO file.
func TwPoToCN(twPOPath, cnPOPath string) {
	os.MkdirAll(filepath.Dir(cnPOPath), 0755)
	fo, err := os.Create(cnPOPath)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	lines, err := util.ReadlinesFromFile(twPOPath)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "msgstr") {
			fo.Write([]byte(util.T2S(line)))
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
