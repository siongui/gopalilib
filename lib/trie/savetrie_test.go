package trie

import (
	"io/ioutil"
	"testing"
)

func TestSaveTrie(t *testing.T) {
	paliwords := []string{"sacca", "ariya", "saccavācā", "dhammaṃ", "buddho", "viharati"}

	b, err := BuildPaliTrieData(paliwords)
	if err != nil {
		t.Error(err)
		return
	}

	err = ioutil.WriteFile("/tmp/trie.json", b, 0644)
	if err != nil {
		t.Error(err)
		return
	}
}
