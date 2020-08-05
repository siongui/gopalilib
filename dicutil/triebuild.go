package dicutil

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/siongui/go-succinct-data-structure-trie"
	"github.com/siongui/gopalilib/util"
)

func BuildSuccinctTrieFromDir(wordsJsonDir, trieDataPath, trieNodeCountPath, rankDirectoryDataPath string) (err error) {
	files, err := ioutil.ReadDir(wordsJsonDir)
	if err != nil {
		return
	}

	var words []string
	for _, file := range files {
		if file.Mode().IsRegular() {
			word := strings.TrimSuffix(file.Name(), ".json")
			words = append(words, word)
		}
	}

	return BuildSuccinctTrie(words, trieDataPath, trieNodeCountPath, rankDirectoryDataPath)
}

func BuildSuccinctTrie(words []string, trieDataPath, trieNodeCountPath, rankDirectoryDataPath string) (err error) {
	// set alphabet of words
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	// encode: build succinct trie
	te := bits.Trie{}
	te.Init()

	for i, word := range words {
		util.LocalPrintln(i+1, word)
		// encode: insert words
		te.Insert(word)
	}

	// encode: trie encoding
	teData := te.Encode()
	//println(teData)
	err = ioutil.WriteFile(trieDataPath, []byte(teData), 0644)
	if err != nil {
		return
	}
	println(te.GetNodeCount())
	err = ioutil.WriteFile(trieNodeCountPath, []byte(strconv.Itoa(int(te.GetNodeCount()))), 0644)
	if err != nil {
		return
	}

	// encode: build cache for quick lookup
	rd := bits.CreateRankDirectory(teData, te.GetNodeCount()*2+1, bits.L1, bits.L2)
	//println(rd.GetData())
	err = ioutil.WriteFile(rankDirectoryDataPath, []byte(rd.GetData()), 0644)
	return
}
