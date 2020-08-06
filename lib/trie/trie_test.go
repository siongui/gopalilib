package trie

import (
	"testing"
)

func TestTrieMethods(t *testing.T) {
	paliwords := []string{"sacca", "ariya", "saccavācā", "dhammaṃ", "buddho", "viharati"}

	b, err := BuildPaliTrieData(paliwords)
	if err != nil {
		t.Error(err)
		return
	}

	ft, err := LoadPaliTrieData(b)
	if err != nil {
		t.Error(err)
		return
	}

	if ft.Lookup("sacca") != true {
		t.Error(`ft.Lookup("sacca")`)
		return
	}

	if ft.Lookup("sacc") == true {
		t.Error(`ft.Lookup("sacc")`)
		return
	}

	if ft.Lookup("saccavācā") != true {
		t.Error(`ft.Lookup("saccavācā")`)
		return
	}

	sw := ft.GetSuggestedWords("d", 10)
	if len(sw) != 1 {
		t.Error("len(sw) != 1")
	}
	if sw[0] != "dhammaṃ" {
		t.Error(`sw[0] != "dhammaṃ"`)
	}
}
