package gettext

import (
	"flag"
	"io/ioutil"
	"testing"
)

var localeDir = flag.String("localeDir", ".", "locale dir containing PO files")

func TestGetPOPath(t *testing.T) {
	if GetPOPath("zh_TW", "messages", "locale") != "locale/zh_TW/LC_MESSAGES/messages.po" {
		t.Error(`GetPOPath("zh_TW", "messages", "locale")`)
		return
	}
}

func TestExtractFromPOFileBytes(t *testing.T) {
	p := GetPOPath("zh_TW", "messages", *localeDir)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		t.Error(err)
		return
	}

	pairs, err := ExtractFromPOFileBytes(b)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(pairs)
}
