package dicmgr

import (
	"testing"
)

func TestLookup(t *testing.T) {
	if Lookup("sacca") != true {
		t.Error(`Lookup("sacca")`)
		return
	}

	if Lookup("sacc") == true {
		t.Error(`Lookup("sacc")`)
		return
	}
}

func TestGetSuggestedWords(t *testing.T) {
	sw := GetSuggestedWords("s", 10)
	if len(sw) != 11 {
		t.Error(`len(sw) != 10`)
		t.Error(sw)
		return
	}
}
