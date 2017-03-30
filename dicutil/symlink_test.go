package dicutil

import (
	"testing"
)

func TestGensite(t *testing.T) {
	err := SymlinkToRootIndexHtml("../../pali/go/website/json/", "../../pali/go/website/")
	if err != nil {
		t.Error(err)
	}
}
