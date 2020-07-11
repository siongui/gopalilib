package lib

import (
	"strings"
	"unicode/utf8"
)

// Type of the webpage, determined according to path of URL
//go:generate stringer -type=PageType
type PageType int

const (
	RootPage PageType = iota
	AboutPage
	PrefixPage
	WordPage
	NoSuchPage
)

// DeterminePageType determines the type of the webpage according to path of
// URL.
func DeterminePageType(urlpath string) PageType {
	if urlpath == "/" {
		return RootPage
	}
	if urlpath == "/about/" {
		return AboutPage
	}
	if IsValidPrefixUrlPath(urlpath) {
		return PrefixPage
	}
	if IsValidWordUrlPath(urlpath) {
		return WordPage
	}

	return NoSuchPage
}

// IsValidPrefixUrlPath will return true for url path such as '/browse/s/' or
// '/browse/ā/'
func IsValidPrefixUrlPath(urlpath string) bool {
	ss := strings.Split(urlpath, "/")

	if len(ss) != 4 {
		return false
	}

	if ss[0] != "" {
		return false
	}

	if ss[1] != "browse" {
		return false
	}

	if ss[3] != "" {
		return false
	}

	if ss[2] != GetFirstCharacterOfWord(ss[2]) {
		return false
	}

	return true
}

// IsValidWordUrlPath will return true if the path of the url is a possible
// valid Pāli word.
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

// If urlpath is valid, return the word string embedded in the path of url.
// Otherwise return empty string.
//
// For example,
//
// "/browse/s/sacca/" will return "sacca"
//
// "/browse/s/āpadā/" will return ""
func GetWordFromUrlPath(urlpath string) string {
	if IsValidWordUrlPath(urlpath) {
		ss := strings.Split(urlpath, "/")
		return ss[3]
	}

	return ""
}

// WordUrlPath will return the url path of the given Pāli word.
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
	return "/browse/" + GetFirstCharacterOfWord(word) + "/" + word + "/"
}

// GetFirstCharacterOfWord returns first character of the word. For example,
// āpadā will return ā
func GetFirstCharacterOfWord(word string) string {
	runeValue, _ := utf8.DecodeRuneInString(word[0:])
	return string(runeValue)
}
