package dicutil

import (
	"encoding/json"
	"github.com/siongui/gopalilib/lib"
	"os"
)

// Return the json file path of word definition.
func GetWordPath(word, wordsJsonDir string) string {
	return wordsJsonDir + "/" + word + ".json"
}

// Given the word and the directory containing definitions of words, return
// definition like the following:
//
//   [
//      [
//         "C",
//         "sukhada\uff1a[adj\uff0e] producing happiness\uff0e"
//      ],
//      [
//         "R",
//         "sukhada\uff1a\u101e\u102f\u1001-\u1012\t\uff08\u1010\u102d\uff09 <br>\u1001\u103a\u1019\u1039\u1038\u101e\u102c\u1000\u102d\u102f \u1031\u1015\u1038\u101e\u100a\u1039\u104b "
//      ]
//   ]
func GetBookIdWordExps(word, wordsJsonDir string) lib.BookIdWordExps {
	f, err := os.Open(GetWordPath(word, wordsJsonDir))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	w := lib.BookIdWordExps{}
	if err := dec.Decode(&w); err != nil {
		panic(err)
	}
	return w
}
