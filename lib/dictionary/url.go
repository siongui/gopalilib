// Package dictionary provides dictionary-specific methods for Pāli Dictionary.
package dictionary

import (
	"net/url"
	"strings"
	"unicode/utf8"
)

// The url of Pali dictionary website except about page will be
//
//   [rootPath]/[locale]/browse/[prefix]/[word]
//
// The about page (all locales share one about page) will be
//
//   [rootPath]/about/
//

// PageType represents the type of the webpage, determined according to path of
// URL.
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

// rootPath is not set directly. This is automatically derived from siteurl.
var rootPath = ""

// SetSiteUrl sets the website url. Used to check path validity and create path
// for pages. Default is none. It's ok not to set this. If the root path of
// deployed Pali dictionary website is "/", there is no need to set site url.
// For example, if Pali dictionary website is deployed at
// "https://dictionary.sutta.org/", there is no need to call this method to set
// site url. However, if root path of deployed Pali dictionary website is other
// than "/", This method must be called during the initialization/loading phase
// of website such that the links in the website can be correctly set. For
// example, if Pali dictionary website is deployed at
// "https://siongui.gitlab.io/pali-dictionary/", it is a must to set site url
// via this method during website initialization/loading.
func SetSiteUrl(rawurl string) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return
	}
	siteurl = rawurl
	rootPath = u.Path
	if strings.HasSuffix(rootPath, "/") {
		// to make process url path more easily later
		rootPath = strings.TrimSuffix(rootPath, "/")
	}
}

// SetCurrentLocale sets the current locale of the website. Used to check path
// validity and create path for pages. Default is none. It's ok not to set this.
func SetCurrentLocale(locale string) {
	currentLocale = locale
}

// The url of Pali dictionary website except about page will be
//
//   [rootPath]/[locale]/browse/[prefix]/[word]
//
// The about page (all locales share one about page) will be
//
//   [rootPath]/about/
//
// This method strip rootPath and locale, and return a "normalized" path for
// furthur processing.
func StripRootPathAndCurrentLocaleInUrlPath(urlpath string) string {
	if len(currentLocale) > 0 {
		ss := strings.Split(urlpath, currentLocale)
		if len(ss) == 2 {
			return ss[1]
		}
	}

	if len(rootPath) > 0 {
		return strings.TrimPrefix(urlpath, rootPath)
	}

	return urlpath
}

// The url of Pali dictionary website except about page will be
//
//   [rootPath]/[locale]/browse/[prefix]/[word]
//
// The about page will be
//
//   [rootPath]/about/
//
// This method add rootPath and locale to /browse/[prefix]/[word]
func AddRootPathAndCurrentLocaleToUrlPath(urlpath string) string {
	if len(currentLocale) > 0 {
		urlpath = "/" + currentLocale + urlpath
	}

	if len(rootPath) > 0 {
		urlpath = rootPath + urlpath
	}

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
