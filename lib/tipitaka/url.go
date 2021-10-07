package tipitaka

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/tipitaka/toc"
)

var actionUrlMap map[string]string
var urlActionMap map[string]string

func traverse(tree lib.Tree, indent int) {
	//print(strings.Repeat(" ", indent))
	//println(TrimTreeText(tree.Text))
	if tree.Action != "" {
		//println(tree.Action)
		actionUrlMap[tree.Action] = ActionToUrlPath(tree.Action)
		urlActionMap[ActionToUrlPath(tree.Action)] = tree.Action
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

// ActionToUrlPath converts action string to url path.
func ActionToUrlPath(action string) string {
	noext := strings.TrimSuffix(action, filepath.Ext(action))
	return "/" + strings.Replace(noext, ".", "/", -1) + "/"
}
