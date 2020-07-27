package util

import (
	"encoding/xml"
	"os"
	"testing"
)

type Tree struct {
	XMLName xml.Name `xml:"tree"`
	Trees   []Tree   `xml:"tree"`
	Text    string   `xml:"text,attr"`
	Src     string   `xml:"src,attr"`
	Action  string   `xml:"action,attr"`
}

func TestDecodeUtf16XML(t *testing.T) {
	dst := "/tmp/romn/cscd/vin01m.mul.toc.xml"
	err := CheckDownload("https://www.tipitaka.org/romn/cscd/vin01m.mul.toc.xml", dst, false)
	if err != nil {
		t.Error(err)
		return
	}

	f16, err := os.Open(dst)
	if err != nil {
		t.Error(err)
		return
	}

	tree := Tree{}
	err = DecodeUtf16XML(f16, &tree)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tree)
}
