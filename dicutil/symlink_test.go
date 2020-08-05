package dicutil

import (
	"flag"
	"testing"

	"pali/words/vfspkg"
)

var outputDir = flag.String("outputDir", ".", "output dir")

func TestSymlinkToRootIndexHtml(t *testing.T) {
	err := SymlinkToRootIndexHtml(*outputDir, vfspkg.MapKeys())
	if err != nil {
		t.Error(err)
	}
}
