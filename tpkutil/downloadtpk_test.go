package tpkutil

import (
	"testing"
)

func TestDownloadTipitaka(t *testing.T) {
	err := DownloadTipitaka("/tmp/tpkxml/", false)
	if err != nil {
		t.Error(err)
		return
	}
}
