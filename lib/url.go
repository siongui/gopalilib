package lib

import (
	"unicode/utf8"
)

// URL path of the Pāli word
func WordUrlPath(word string) string {
	runeValue, _ := utf8.DecodeRuneInString(word[0:])
	return "/browse/" + string(runeValue) + "/" + word + "/"
}
