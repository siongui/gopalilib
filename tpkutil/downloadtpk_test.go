package tpkutil

import (
	"testing"
)

func TestDownloadTipitaka(t *testing.T) {
	err := DownloadTipitaka("/tmp/", false)
	if err != nil {
		t.Error(err)
		return
	}
}
