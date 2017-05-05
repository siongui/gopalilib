package dicutil

import (
	"testing"
)

func TestSymlinkToRootIndexHtml(t *testing.T) {
	err := SymlinkToRootIndexHtml("../../pali/go/website/json/", "../../pali/go/website/")
	if err != nil {
		t.Error(err)
	}
}
