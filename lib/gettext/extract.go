package gettext

import (
	"path"
	"regexp"
)

const pattern = `msgid "(.+)"\nmsgstr "(.+)"`

// GetPOPath returns the conventional path of PO file.
func GetPOPath(locale, domain, localedir string) string {
	filename := domain + ".po"
	return path.Join(localedir, locale, "LC_MESSAGES", filename)
}

// ExtractFromPOFileBytes returns pairs of (msgid, msgstr), given the bytes of
// the PO file. The bytes of PO file can be read using ioutil.ReadFile in Go
// standard library.
func ExtractFromPOFileBytes(b []byte) (pairs MsgIdStrPairs, err error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return
	}
	matches := re.FindAllStringSubmatch(string(b), -1)

	pairs = MsgIdStrPairs{}
	for _, array := range matches {
		pairs[array[1]] = array[2]
	}
	return
}
