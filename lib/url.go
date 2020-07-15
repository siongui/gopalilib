package lib

import (
	"net/url"
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

var siteurl = ""
var currentLocale = ""

// SetSiteUrl sets the website url. Used to check path validity and create path
// for pages. Default is none. It's ok not to set this.
func SetSiteUrl(u string) {
	siteurl = u
}

// SetCurrentLocale sets the current locale of the website. Used to check path
// validity and create path for pages. Default is none. It's ok not to set this.
func SetCurrentLocale(locale string) {
	currentLocale = locale
}

func StripRootPathAndCurrentLocaleInUrlPath(urlpath string) string {
	if len(siteurl) > 0 {
		// to be implemented
	}

	if len(currentLocale) > 0 {
		// to be implemented
	}

	return urlpath
}

func AddRootPathAndCurrentLocaleToUrlPath(urlpath string) string {
	// to be implemented
	return urlpath
}

// DeterminePageType determines the type of the webpage according to path of
// URL.
func DeterminePageType(urlpath string) PageType {
	// handle url.PathUnescape error?
	urlpath, _ = url.PathUnescape(urlpath)
	urlpath = StripRootPathAndCurrentLocaleInUrlPath(urlpath)

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

// IsValidPrefixUrlPath will return true if the path of the url is a possible
// prefix of Pāli words.
func IsValidPrefixUrlPath(urlpath string) bool {
	// handle url.PathUnescape error?
	urlpath, _ = url.PathUnescape(urlpath)
	urlpath = StripRootPathAndCurrentLocaleInUrlPath(urlpath)

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

// IsValidWordUrlPath will return true if the path of the url is a possible Pāli
// word.
func IsValidWordUrlPath(urlpath string) bool {
	// handle url.PathUnescape error?
	urlpath, _ = url.PathUnescape(urlpath)
	urlpath = StripRootPathAndCurrentLocaleInUrlPath(urlpath)

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

// GetPrefixFromUrlPath will return the prefix string embedded in the path of
// url if url path is valid. Otherwise return empty string. Note that this
// method do not check if the prefix string is a valid prefix. Use with caution.
//
// For example,
//
// "/browse/s/" will return "s"
//
// "/browse/āā/" will return ""
func GetPrefixFromUrlPath(urlpath string) string {
	// handle url.PathUnescape error?
	urlpath, _ = url.PathUnescape(urlpath)
	urlpath = StripRootPathAndCurrentLocaleInUrlPath(urlpath)

	if IsValidPrefixUrlPath(urlpath) {
		ss := strings.Split(urlpath, "/")
		return ss[2]
	}

	return ""
}

// GetWordFromUrlPath will return the word string embedded in the path of url if
// url path is valid. Otherwise return empty string. Note that this method do
// not check if the word string is a valid word. Use with caution.
//
// For example,
//
// "/browse/s/sacca/" will return "sacca"
//
// "/browse/s/āpadā/" will return ""
func GetWordFromUrlPath(urlpath string) string {
	// handle url.PathUnescape error?
	urlpath, _ = url.PathUnescape(urlpath)
	urlpath = StripRootPathAndCurrentLocaleInUrlPath(urlpath)

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
// Note that this method do not check the validity of the word. Use with
// caution.
func WordUrlPath(word string) string {
	urlpath := "/browse/" + GetFirstCharacterOfWord(word) + "/" + word + "/"
	urlpath = AddRootPathAndCurrentLocaleToUrlPath(urlpath)
	return urlpath
}

// GetFirstCharacterOfWord returns first character of the word. For example,
// āpadā will return ā
func GetFirstCharacterOfWord(word string) string {
	runeValue, _ := utf8.DecodeRuneInString(word)
	return string(runeValue)
}

// PrefixUrlPath will return the url path of the given prefix.
//
// Example:
//
// URL path of prefix ``s`` is:
//
//   /browse/s/
//
// URL path of prefix ``ā`` is:
//
//   /browse/ā/
//
// Note that this method do not check the validity of the prefix. Use with
// caution.
func PrefixUrlPath(prefix string) string {
	urlpath := "/browse/" + prefix + "/"
	urlpath = AddRootPathAndCurrentLocaleToUrlPath(urlpath)
	return urlpath
}
