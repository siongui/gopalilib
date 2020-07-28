package tpkutil

import (
	"testing"
)

func TestBuildTipitakaTree(t *testing.T) {
	err := BuildTipitakaTree("/tmp/tpkxml/")
	if err != nil {
		t.Error(err)
		return
	}
}
