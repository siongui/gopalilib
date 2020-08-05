package dicutil

import (
	"flag"
	"testing"
)

var wordsJsonDir = flag.String("wordsJsonDir", ".", "output dir of json files of pali words")
var trieData = flag.String("trieData", ".", "txt file of trie data")
var trieNodeCount = flag.String("trieNodeCount", ".", "txt file of trie node count")
var trieRankDirectoryData = flag.String("trieRankDirectoryData", ".", "txt file of trie rank directory data")

func TestBuildSuccinctTrie(t *testing.T) {
	BuildSuccinctTrie(*wordsJsonDir, *trieData, *trieNodeCount, *trieRankDirectoryData)
}
