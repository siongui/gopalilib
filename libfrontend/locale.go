package libfrontend

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib/jsgettext"
)

func GetFinalShowLocale() string {
	var supportedLocales = []string{"en_US", "zh_TW", "vi_VN", "fr_FR"}
	var navigatorLanguages = Window.Navigator().Languages()
	// show language according to site url and NavigatorLanguages API
	locale := Document.GetElementById("site-info").Dataset().Get("locale").String()
	if locale == "" {
		return jsgettext.DetermineLocaleByNavigatorLanguages(navigatorLanguages, supportedLocales)
	}
	return locale
}

func TranslateDocument(locale string) {
	elms := Document.QuerySelectorAll("[data-default-string]")
	for _, elm := range elms {
		str := elm.Get("dataset").Get("defaultString").String()
		elm.Set("textContent", jsgettext.Gettext(locale, str))
	}
}
