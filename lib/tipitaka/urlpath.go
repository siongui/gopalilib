// Package tipitaka provides tipitaka-specific methods for Pāli Tipiṭaka.
// This package consists of common data structures and methods to be shared and
// used in front-end (browser), back-end (server), and offline data processing.
package tipitaka

import (
	"regexp"
	"strings"

	"github.com/siongui/gopalilib/lib"
)

// remove leading un-needed characters
var tp = regexp.MustCompile(`^[\d\s()-\.]+`)

// remove trailing un-needed characters
var ts = regexp.MustCompile(`-\d$`)

// TrimTreeText trims Text property in Tree struct. The same as old Python
// implementation.
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

func TrimTreeText2(text string) string {
	text = TrimTreeText(text)
	// not in old Python implementation
	text = strings.TrimSuffix(text, "kathā")
	text = strings.TrimSuffix(text, "pariccheda")
	text = strings.TrimSuffix(text, "paricchedo")
	return strings.TrimSpace(text)
}

func traverse(tree lib.Tree, indent int) {
	//print(strings.Repeat(" ", indent))
	//println(TrimTreeText(tree.Text))
	for _, subtree := range tree.SubTrees {
		traverse(subtree, indent+2)
	}
}

func TraverseTreeAndSetSubpathProperty(tree lib.Tree) {
	traverse(tree, 0)
}
