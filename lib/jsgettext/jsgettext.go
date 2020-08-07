package jsgettext

import (
	"github.com/siongui/gopalilib/lib/gettext"
)

func init() {
	b, err := ReadFile("po.json")
	if err != nil {
		panic(err)
	}

	err = gettext.SetupTranslationMapping(b)
	if err != nil {
		panic(err)
	}
}

func Gettext(locale, str string) string {
	return gettext.Gettext(locale, str)
}

func DetermineLocaleByNavigatorLanguages(languages string, supportedLocales []string) string {
	return gettext.DetermineLocaleByNavigatorLanguages(languages, supportedLocales)
}
