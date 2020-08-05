package dicutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/siongui/go-succinct-data-structure-trie"
	"github.com/siongui/gopalilib/util"
)

func BuildSuccinctTrie(wordsJsonDir, trieDataPath, trieNodeCountPath, rankDirectoryDataPath string) {
	// set alphabet of words
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	// encode: build succinct trie
	te := bits.Trie{}
	te.Init()

	i := 0
	// walk all word json files
	filepath.Walk(wordsJsonDir, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			word := strings.TrimSuffix(info.Name(), ".json")
			util.LocalPrintln(i, word)
			// encode: insert words
			te.Insert(word)
			i++
		}
		return nil
	})

	// encode: trie encoding
	teData := te.Encode()
	//println(teData)
	ioutil.WriteFile(trieDataPath, []byte(teData), 0644)
	println(te.GetNodeCount())
	ioutil.WriteFile(trieNodeCountPath, []byte(strconv.Itoa(int(te.GetNodeCount()))), 0644)

	// encode: build cache for quick lookup
	rd := bits.CreateRankDirectory(teData, te.GetNodeCount()*2+1, bits.L1, bits.L2)
	//println(rd.GetData())
	ioutil.WriteFile(rankDirectoryDataPath, []byte(rd.GetData()), 0644)
}
