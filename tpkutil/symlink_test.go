package tpkutil

import (
	"flag"
	"testing"
)

var outputDir = flag.String("outputDir", ".", "output dir")

func TestSymlinkToRootIndexHtml(t *testing.T) {
	err := SymlinkToRootIndexHtml(*outputDir, "romn")
	if err != nil {
		t.Error(err)
	}
}
