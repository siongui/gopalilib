package tipitaka

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dictionary"
	"github.com/siongui/gopalilib/lib/tipitaka/toc"
)

// The url path of Pali tipitaka website will be
//
//   [rootPath]/[locale]/[canonPath]

var actionUrlMap map[string]string
var urlActionMap map[string]string

func traverse(tree lib.Tree, indent int) {
	//print(strings.Repeat(" ", indent))
	//println(TrimTreeText(tree.Text))
	if tree.Action != "" {
		//println(tree.Action)
		actionUrlMap[tree.Action] = ActionToCanonPath(tree.Action)
		urlActionMap[ActionToCanonPath(tree.Action)] = tree.Action
	}
	for _, subtree := range tree.SubTrees {
		traverse(subtree, indent+2)
	}
}

func init() {
	actionUrlMap = make(map[string]string)
	urlActionMap = make(map[string]string)

	b, _ := toc.ReadFile("tpktoc.json")
	//println(string(b))
	tree := lib.Tree{}
	json.Unmarshal(b, &tree)
	traverse(tree, 0)
}

// ActionToCanonPath converts action string to canon path in URL.
func ActionToCanonPath(action string) string {
	noext := strings.TrimSuffix(action, filepath.Ext(action))
	// TODO: FIXME: elegant way to support different script and edition.
	return "/romn/" + strings.Replace(noext, ".", "/", -1) + "/"
}

// GetAllCanonPath returns all canon paths according to given script.
func GetAllCanonPath(script string) []string {
	// FIXME TODO: script param is not respected right now. return only romn

	// https://stackoverflow.com/a/27848197
	keys := make([]string, len(urlActionMap))
	i := 0
	for k := range urlActionMap {
		keys[i] = k
		i++
	}

	return keys
}

// PageType represents the type of the webpage, determined according to path of
// URL.
//go:generate stringer -type=PageType
type PageType int

const (
	RootPage PageType = iota
	CanonPage
	NoSuchPage
)

// SetSiteUrl is the same as the method of the same name in the package of
// github.com/siongui/gopalilib/lib/dictionary
func SetSiteUrl(rawurl string) {
	dictionary.SetSiteUrl(rawurl)
}

// SetCurrentLocale is the same as the method of the same name in the package of
// github.com/siongui/gopalilib/lib/dictionary
func SetCurrentLocale(locale string) {
	dictionary.SetCurrentLocale(locale)
}

// DeterminePageType determines the type of the webpage according to path of
// URL.
func DeterminePageType(urlpath string) PageType {
	urlpath, _ = dictionary.GetNormalizedUrlPath(urlpath)

	if urlpath == "/" {
		return RootPage
	}
	if IsValidCanonUrlPath(urlpath) {
		return CanonPage
	}

	return NoSuchPage
}

// IsValidCanonUrlPath will return true if the path of the url is a possible
// canon page.
func IsValidCanonUrlPath(urlpath string) bool {
	urlpath, _ = dictionary.GetNormalizedUrlPath(urlpath)

	_, ok := urlActionMap[urlpath]
	return ok
}

// ActionToUrlPath converts action string to url path.
func ActionToUrlPath(action string) string {
	return dictionary.AddRootPathAndCurrentLocaleToUrlPath(ActionToCanonPath(action))
}
