package tpkutil

import (
	"testing"
)

func TestDownloadTipitaka(t *testing.T) {
	//err := DownloadTipitaka("romn", "/tmp/tpkxml/", false)
	err := GetAllXml("https://tipitaka.org/romn/", "cscd/vin01m.mul.toc.xml", "/tmp/tpkxml/", false)
	if err != nil {
		t.Error(err)
		return
	}
}
