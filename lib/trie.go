package lib

import (
	"encoding/json"

	bits "github.com/siongui/go-succinct-data-structure-trie"
)

type TrieData struct {
	EncodedData       string
	NodeCount         uint
	RankDirectoryData string
}

// TODO: add func to set allowed characters

// TrieToJson outputs JSON []byte of TrieData struct.
func TrieToJson(t bits.Trie) ([]byte, error) {
	// encode: build cache for quick lookup
	rd := bits.CreateRankDirectory(t.Encode(), t.GetNodeCount()*2+1, bits.L1, bits.L2)
	td := TrieData{
		// encode: trie encoding
		EncodedData:       t.Encode(),
		NodeCount:         t.GetNodeCount(),
		RankDirectoryData: rd.GetData(),
	}

	return json.Marshal(td)
}

// JsonToTrie loads TrieData struct from JSON []byte.
func JsonToTrie(b []byte) (td TrieData, err error) {
	err = json.Unmarshal(b, &td)
	return
}

// BuildPaliTrieData outputs JSON format trie data.
func BuildPaliTrieData(paliwords []string) ([]byte, error) {
	// set alphabet of words
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	// encode: build succinct trie
	t := bits.Trie{}
	t.Init()

	for _, word := range paliwords {
		// encode: insert word
		t.Insert(word)
	}

	return TrieToJson(t)
}

// LoadPaliTrieData returns frozen trie from JSON []byte of TrieData struct.
func LoadPaliTrieData(b []byte) (ft bits.FrozenTrie, err error) {
	// set alphabet of words
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")

	td, err := JsonToTrie(b)
	if err != nil {
		return
	}

	// decode: build frozen succinct trie
	ft.Init(td.EncodedData, td.RankDirectoryData, td.NodeCount)
	return
}