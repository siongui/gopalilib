package lib

import (
	"strings"
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
	if IsValidWordUrlPath(urlpath) {
		return WordPage
	}

	return NoSuchPage
}

// Give the path of url, is it a possible valid path for a Pāli word?
func IsValidWordUrlPath(urlpath string) bool {
	ss := strings.Split(urlpath, "/")

	if len(ss) != 5 {
		return false
	}

	if ss[0] != "" {
		return false
	}

	if ss[1] != "browse" {
		return false
	}

	if ss[4] != "" {
		return false
	}

	if !strings.HasPrefix(ss[3], ss[2]) {
		return false
	}

	return true
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
