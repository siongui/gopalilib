package dicutil

import (
	"flag"
	"testing"

	vfs "github.com/siongui/gopaliwordvfs"
)

var outputDir = flag.String("outputDir", ".", "output dir")

func TestSymlinkToRootIndexHtml(t *testing.T) {
	err := SymlinkToRootIndexHtml(*outputDir, vfs.MapKeys())
	if err != nil {
		t.Error(err)
	}
}
