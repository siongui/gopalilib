// Package dicmgr provides high-level methods to access dictionary data.
package dicmgr

import (
	"encoding/json"

	bits "github.com/siongui/go-succinct-data-structure-trie"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/trie"
)

var di lib.BookIdAndInfos
var ft bits.FrozenTrie

func init() {
	b, err := ReadFile("BookIdAndInfos.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &di)
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

func Lookup(word string) bool {
	return ft.Lookup(word)
}
