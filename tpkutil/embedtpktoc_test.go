package tpkutil

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/siongui/goef"
	"github.com/siongui/gopalilib/util"
)

var tpkXmlDir = flag.String("tpkXmlDir", ".", "xml dir of pali tipitaka")
var outputGoFilePath = flag.String("outputGoFilePath", "lib/tipitaka/toc/data.go", "Go file containing data")

func TestCreateTipitakaTableOfContentGoCode(t *testing.T) {
	tree, err := BuildTipitakaTree(*tpkXmlDir)
	if err != nil {
		t.Error(err)
		return
	}

	b, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
		return
	}

	tmpdir, err := ioutil.TempDir("", "tpk-toc-tmp-dir")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(tmpdir) // clean up

	err = ioutil.WriteFile(filepath.Join(tmpdir, "tpktoc.json"), b, 0644)
	if err != nil {
		t.Error(err)
		return
	}

	util.CreateDirIfNotExist(*outputGoFilePath)
	err = goef.GenerateGoPackagePlainText("toc", tmpdir, *outputGoFilePath)
	if err != nil {
		t.Error(err)
		return
	}
}
