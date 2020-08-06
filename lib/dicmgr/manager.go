// Package dicmgr provides high-level methods to access dictionary data.
// This package consists of common data structures and methods to be shared and
// used in front-end (browser), back-end (server), and offline data processing.
package dicmgr

import (
	"encoding/json"

	bits "github.com/siongui/go-succinct-data-structure-trie"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/trie"
)

var bookIdAndInfos lib.BookIdAndInfos
var ft bits.FrozenTrie

func init() {
	b, err := ReadFile("BookIdAndInfos.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &bookIdAndInfos)
	if err != nil {
		panic(err)
	}

	b, err = ReadFile("trie.json")
	if err != nil {
		panic(err)
	}

	ft, err = trie.LoadPaliTrieData(b)
	if err != nil {
		panic(err)
	}
}

// Lookup returns if the word is in the dictionary.
func Lookup(word string) bool {
	return ft.Lookup(word)
}

// GetSuggestedWords returns suggested words starting with prefix, and the
// number of returned suggested words is limited by second argument.
func GetSuggestedWords(prefix string, limit int) []string {
	return ft.GetSuggestedWords(prefix, limit)
}
