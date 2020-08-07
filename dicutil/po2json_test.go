package dicutil

import (
	"flag"
	"testing"

	"github.com/siongui/gopalilib/lib/gettext"
)

var localeDir = flag.String("localeDir", ".", "locale dir containing PO files")
var outputGoDataFilePath = flag.String("outputGoDataFilePath", "lib/gettext/data.go", "Go file containing data")

func TestGettext(t *testing.T) {
	b, err := PO2JSONBytes("messages", *localeDir)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(b))

	err = gettext.SetupTranslationMapping(b)
	if err != nil {
		t.Error(err)
		return
	}

	s := gettext.Gettext("zh_TW", "Pāli Dictionary")
	if s != "巴利字典" {
		t.Error(s)
		return
	}
}

func TestEmbedPOJSONInGoCode(t *testing.T) {
	err := EmbedPOJSONInGoCode("messages", *localeDir, "gettext", *outputGoDataFilePath)
	if err != nil {
		t.Error(err)
		return
	}
}
