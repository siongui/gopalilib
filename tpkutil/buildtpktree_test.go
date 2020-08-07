package tpkutil

import (
	"flag"
	"testing"
)

var tpkXmlDir = flag.String("tpkXmlDir", ".", "xml dir of pali tipitaka")

func TestBuildTipitakaTree(t *testing.T) {
	tree, err := BuildTipitakaTree(*tpkXmlDir)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tree)
}
