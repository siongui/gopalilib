package dicutil

import (
	"flag"
	"path"
	"testing"
)

var pkgdir = flag.String("pkgdir", ".", "dir of package containing embedded files")

func TestBuildVFS(t *testing.T) {
	pkgName := path.Base(*pkgdir)
	t.Log("pkgName: ", pkgName)
	t.Log("outputDir: ", *pkgdir)

	err := BuildVFS(pkgName, wordsJsonDir, *pkgdir)
	if err != nil {
		t.Error(err)
	}
}
