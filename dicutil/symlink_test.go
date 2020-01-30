package dicutil

import (
	"testing"
)

func TestSymlinkToRootIndexHtml(t *testing.T) {
	err := SymlinkToRootIndexHtml(wordsJsonDir)
	if err != nil {
		t.Error(err)
	}
}
