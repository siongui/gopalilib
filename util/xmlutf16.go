package util

import (
	"encoding/xml"
	"io"

	"golang.org/x/net/html/charset"
)

func BypassReader(label string, input io.Reader) (io.Reader, error) {
	return input, nil
}

func DecodeUtf16XML(r io.Reader, v interface{}) (err error) {
	// Fail to decode: https://www.tipitaka.org/romn/cscd/vin01m.mul.toc.xml
	// The Tipiṭaka XML is encoded in UTF-16
	// Google search: golang xml utf-16
	// https://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go
	// https://groups.google.com/forum/#!topic/golang-nuts/tXcECEKC2rs
	nr, err := charset.NewReader(r, "utf-16")
	if err != nil {
		return
	}
	decoder := xml.NewDecoder(nr)
	decoder.CharsetReader = BypassReader
	err = decoder.Decode(v)
	return
}
