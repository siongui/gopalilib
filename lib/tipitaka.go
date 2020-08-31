package lib

import (
	"encoding/xml"
)

// Tree struct represents Table of Content (ToC) node of Tipiṭaka
type Tree struct {
	XMLName  xml.Name `xml:"tree"`
	SubTrees []Tree   `xml:"tree"`
	Text     string   `xml:"text,attr"`
	Src      string   `xml:"src,attr"`
	Action   string   `xml:"action,attr"`
}
