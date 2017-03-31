package lib

import (
	"unicode/utf8"
)

// Type of the webpage, determined according to path of URL
type PageType int

const (
	RootPage = iota
	AboutPage
	WordPage
	NoSuchPage
)

// Determine the type of the webpage according to path of URL.
func DeterminePageType(urlpath string) PageType {
	if urlpath == "/" {
		return RootPage
	}
	if urlpath == "/about/" {
		return AboutPage
	}

	return WordPage
}

// URL path of the Pāli word.
//
// Example:
//
// URL path of word ``sacca`` is:
//
//   /browse/s/sacca/
//
// URL path of word ``āpadā`` is:
//
//   /browse/ā/āpadā/
//
func WordUrlPath(word string) string {
	runeValue, _ := utf8.DecodeRuneInString(word[0:])
	return "/browse/" + string(runeValue) + "/" + word + "/"
}
