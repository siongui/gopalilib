package dicutil

import (
	"io/ioutil"
	"strings"

	"github.com/siongui/gopalilib/lib/trie"
	"github.com/siongui/gopalilib/util"
)

func BuildSuccinctTrieFromDir(wordsJsonDir, trieJsonPath string) (err error) {
	files, err := ioutil.ReadDir(wordsJsonDir)
	if err != nil {
		return
	}

	var words []string
	for i, file := range files {
		if file.Mode().IsRegular() {
			word := strings.TrimSuffix(file.Name(), ".json")
			words = append(words, word)
			util.LocalPrintln(i, word)
		}
	}

	b, err := trie.BuildPaliTrieData(words)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(trieJsonPath, b, 0644)
	return
}
