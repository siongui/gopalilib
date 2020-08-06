package dicutil

import (
	"flag"
	"testing"
)

var wordsJsonDir = flag.String("wordsJsonDir", ".", "output dir of json files of pali words")
var trieJson = flag.String("trieJson", "/tmp/trie.json", "json file of trie data")

func TestBuildSuccinctTrieFromDir(t *testing.T) {
	err := BuildSuccinctTrieFromDir(*wordsJsonDir, *trieJson)
	if err != nil {
		t.Error(err)
		return
	}
}
