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
//   [rootPath]/[locale]/[paliTextPath]

var actionToPaliTextPathMap map[string]string
var paliTextPathToActionMap map[string]string

func traverse(tree lib.Tree, indent int) {
	//print(strings.Repeat(" ", indent))
	//println(TrimTreeText(tree.Text))
	if tree.Action != "" {
		//println(tree.Action)
		actionToPaliTextPathMap[tree.Action] = ActionToPaliTextPath(tree.Action)
		paliTextPathToActionMap[ActionToPaliTextPath(tree.Action)] = tree.Action
	}
	for _, subtree := range tree.SubTrees {
		traverse(subtree, indent+2)
	}
}

func init() {
	actionToPaliTextPathMap = make(map[string]string)
	paliTextPathToActionMap = make(map[string]string)

	b, _ := toc.ReadFile("tpktoc.json")
	//println(string(b))
	tree := lib.Tree{}
	json.Unmarshal(b, &tree)
	traverse(tree, 0)
}

// ActionToPaliTextPath converts action string to pali text path in URL.
// For example, "cscd/vin01m.mul2.xml" to "/romn/cscd/vin01m/mul2/".
// TODO FIXME: add *script* and *edition* parameters in the future.
func ActionToPaliTextPath(action string) string {
	noext := strings.TrimSuffix(action, filepath.Ext(action))
	// TODO FIXME: find elegant way to support different script and edition.
	return "/romn/" + strings.Replace(noext, ".", "/", -1) + "/"
}

// GetAllPaliTextPath returns all pali text paths according to given script.
// TODO FIXME: add *edition* parameter in the future.
func GetAllPaliTextPath(script string) []string {
	// FIXME TODO: script param is not respected right now. return only romn

	// https://stackoverflow.com/a/27848197
	keys := make([]string, len(paliTextPathToActionMap))
	i := 0
	for k := range paliTextPathToActionMap {
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
	PaliTextPage
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
	if ok, _ := IsValidPaliTextUrlPath(urlpath); ok {
		return PaliTextPage
	}

	return NoSuchPage
}

// IsValidPaliTextUrlPath will return both true and pali text path if the path
// of the url is a possible pali text page.
func IsValidPaliTextUrlPath(urlpath string) (ok bool, paliTextPath string) {
	paliTextPath, _ = dictionary.GetNormalizedUrlPath(urlpath)

	_, ok = paliTextPathToActionMap[paliTextPath]
	return
}

// ActionToUrlPath converts action string to url path.
func ActionToUrlPath(action string) string {
	return dictionary.AddRootPathAndCurrentLocaleToUrlPath(ActionToPaliTextPath(action))
}
