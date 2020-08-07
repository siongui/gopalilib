package gettext

import (
	"strings"
)

func DetermineLocaleByNavigatorLanguages(languages string, supportedLocales []string) string {
	for _, language := range strings.Split(languages, ",") {
		language = strings.TrimSpace(language)
		for _, locale := range supportedLocales {
			localeNormalized := strings.Replace(locale, "_", "-", -1)
			if language == localeNormalized {
				return locale
			}
			if language[:2] == locale[:2] {
				return locale
			}
		}
	}
	return supportedLocales[0]
}
