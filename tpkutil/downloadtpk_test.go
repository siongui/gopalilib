package tpkutil

import (
	"testing"
)

func TestDownloadTipitaka(t *testing.T) {
	err := DownloadTipitaka("/tmp/")
	if err != nil {
		t.Error(err)
		return
	}
}
