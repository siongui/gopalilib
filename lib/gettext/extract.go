package gettext

import (
	"path"
	"regexp"
)

const pattern = `msgid "(.+)"\nmsgstr "(.+)"`

func GetPOPath(locale, domain, localedir string) string {
	filename := domain + ".po"
	return path.Join(localedir, locale, "LC_MESSAGES", filename)
}

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
