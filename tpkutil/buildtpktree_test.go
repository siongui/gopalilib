package tpkutil

import (
	"testing"
)

func TestBuildTipitakaTree(t *testing.T) {
	tree, err := BuildTipitakaTree("/tmp/tpkxml/")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tree)
}
