package lib

import (
	"encoding/xml"
	"regexp"
	"strings"
)

// Tree struct represents Table of Content (ToC) node of Tipiṭaka
type Tree struct {
	XMLName  xml.Name `xml:"tree"`
	SubTrees []Tree   `xml:"tree"`
	Text     string   `xml:"text,attr"`
	Src      string   `xml:"src,attr"`
	Action   string   `xml:"action,attr"`
}

// remove leading un-needed characters
var tp = regexp.MustCompile(`^[\d\s()-\.]+`)

// remove trailing un-needed characters
var ts = regexp.MustCompile(`-\d$`)

// TrimTreeText trims Text property in Tree struct.
//
// TODO: handle path conflict after TrimTreeText
// For example,
// Sāratthadīpanī-ṭīkā-1
// Sāratthadīpanī-ṭīkā-2
// Sāratthadīpanī-ṭīkā-3
// are the same after TrimTreeText
func TrimTreeText(text string) string {
	text = tp.ReplaceAllString(text, "")
	text = ts.ReplaceAllString(text, "")
	text = strings.ToLower(text)
	text = strings.TrimSuffix(text, "pāḷi")
	text = strings.TrimSuffix(text, "nikāya")
	text = strings.TrimSuffix(text, "piṭaka")
	text = strings.TrimSuffix(text, "piṭaka (aṭṭhakathā)")

	text = strings.TrimSuffix(text, "kaṇḍa-aṭṭhakathā")
	text = strings.TrimSuffix(text, "-aṭṭhakathā")

	text = strings.TrimSuffix(text, " nikāya (aṭṭhakathā)")
	text = strings.TrimSuffix(text, "nikāya (aṭṭhakathā)")

	text = strings.TrimSuffix(text, "piṭaka (ṭīkā)")
	text = strings.TrimSuffix(text, "nikāya (ṭīkā)")
	text = strings.TrimSuffix(text, "-mūlaṭīkā")
	text = strings.TrimSuffix(text, "-ṭīkā")
	return text
}
