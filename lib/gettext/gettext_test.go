package gettext

import (
	"encoding/json"
	//	"flag"
	"io/ioutil"
	"testing"
)

//var localeDir = flag.String("localeDir", ".", "locale dir containing PO files")

// PO2JSON converts PO files to JSON bytes.
func PO2JSONBytes(domain, localedir string) (b []byte, err error) {
	dirs, err := ioutil.ReadDir(localedir)
	if err != nil {
		return
	}

	// create PO-like json data for i18n
	obj := LocalesMsg{}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		locale := dir.Name()
		// English is default language
		if locale == "en_US" {
			continue
		}

		b, err = ioutil.ReadFile(GetPOPath(locale, domain, localedir))
		if err != nil {
			return
		}

		obj[locale], err = ExtractFromPOFileBytes(b)
		if err != nil {
			return
		}
	}

	b, err = json.Marshal(obj)
	return
}

func TestGettext(t *testing.T) {
	b, err := PO2JSONBytes("messages", *localeDir)
	if err != nil {
		t.Error(err)
		return
	}

	err = SetupTranslationMapping(b)
	if err != nil {
		t.Error(err)
		return
	}

	s := Gettext("zh_TW", "Pāli Dictionary")
	if s != "巴利字典" {
		t.Error(s)
		return
	}
}
