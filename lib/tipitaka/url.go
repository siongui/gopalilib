package tipitaka

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/siongui/gopalilib/lib"
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
	return "/" + strings.Replace(noext, ".", "/", -1) + "/"
}
