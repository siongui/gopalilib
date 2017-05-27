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
//   {
//     "C": "sukhada：[adj．] producing happiness．",
//     "R": "sukhada：သုခ-ဒ\t（တိ） <br>ခ်မ္းသာကို ေပးသည္။ "
//   }
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
