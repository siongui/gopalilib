package dicutil

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/siongui/goef"
	"github.com/siongui/gopalilib/lib/gettext"
	"github.com/siongui/gopalilib/util"
)

// PO2JSON converts PO files to JSON bytes.
func PO2JSONBytes(domain, localedir string) (b []byte, err error) {
	dirs, err := ioutil.ReadDir(localedir)
	if err != nil {
		return
	}

	// create PO-like json data for i18n
	obj := gettext.LocalesMsg{}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		locale := dir.Name()
		// English is default language
		if locale == "en_US" {
			continue
		}

		b, err = ioutil.ReadFile(gettext.GetPOPath(locale, domain, localedir))
		if err != nil {
			return
		}

		obj[locale], err = gettext.ExtractFromPOFileBytes(b)
		if err != nil {
			return
		}
	}

	b, err = json.Marshal(obj)
	return
}

func EmbedPOJSONInGoCode(domain, localedir, pkgName, outputGoDataFilePath string) (err error) {
	b, err := PO2JSONBytes(domain, localedir)
	if err != nil {
		return
	}

	util.CreateDirIfNotExist(outputGoDataFilePath)
	dir, err := ioutil.TempDir("", "po2jsontmp")
	if err != nil {
		return
	}
	defer os.RemoveAll(dir) // clean up

	err = ioutil.WriteFile(path.Join(dir, "po.json"), b, 0644)
	if err != nil {
		return
	}

	return goef.GenerateGoPackagePlainText(pkgName, dir, outputGoDataFilePath)
}
