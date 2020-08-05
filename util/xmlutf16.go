package util

import (
	"encoding/xml"
	"io"

	"golang.org/x/net/html/charset"
)

// BypassReader is a workaround to process UTF-16 XML. See
//
//   https://groups.google.com/forum/#!topic/golang-nuts/tXcECEKC2rs
//   https://stackoverflow.com/a/50812725
//
// for more information.
func BypassReader(label string, input io.Reader) (io.Reader, error) {
	return input, nil
}

// DecodeUtf16XML decodes UTF-16 encoded XML. The first argument could be an
// UTF-16 encoded XML file opened by os.Open method. The seconde argument is the
// struct used to decode the XML.
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
