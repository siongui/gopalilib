// Package jsgettext is built on top of package gettext (located at lib/gettext
// in the same repo). The only difference is that translation data is directly
// included in this package so that there is no need to call
// SetupTranslationMapping method to setup translation.
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
