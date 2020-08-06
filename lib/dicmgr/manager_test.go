package dicmgr

import (
	"testing"
)

func TestLookup(t *testing.T) {
	if Lookup("sacca") != true {
		t.Error(`Lookup("sacca")`)
		return
	}
}
