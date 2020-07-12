// Package dicuitl provides methods for offline processing to build Pāli
// Dictionary.
package dicutil

// This file make symlinks for Pāli Dictionary SPA (Single Page Application)
// hosted on GitHub Pages.

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/util"
	vfs "github.com/siongui/gopaliwordvfs"
)

func printWord(word string) {
	dontPrintEnv := []string{"TRAVIS", "GITLAB_CI"}
	for _, ev := range dontPrintEnv {
		if _, ok := os.LookupEnv(ev); ok {
			return
		}
	}
	fmt.Println(word)
}

// The URL path of word:
//
//   /browse/{{first char of word}}/{{word}}/
//
func CreateSymlink(word, root string) (err error) {
	// create dir of the word
	wordIndexAbs := filepath.Join(root, lib.WordUrlPath(word), "index.html")
	util.CreateDirIfNotExist(wordIndexAbs)
	//fmt.Println("wordIndexAbs:", wordIndexAbs)

	err = os.Chdir(root)
	if err != nil {
		return
	}

	wordIndex := filepath.Join(lib.WordUrlPath(word), "index.html")
	// remove heading /
	wordIndex = wordIndex[1:]
	//fmt.Println("word index.html path:", wordIndex)
	err = os.Symlink("../../../index.html", wordIndex)
	if os.IsExist(err) {
		// If the symlink we want to create already exist, error will be
		// raised. Remove the existing symlink amd create new symlink.
		os.Remove(wordIndex)
		err = os.Symlink("../../../index.html", wordIndex)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}

	printWord(word)
	return
}

// Only one page: ``/index.html``
//
// All other webpages are symlinks to ``/index.html``
func SymlinkToRootIndexHtml(websiteroot string) (err error) {
	websiteroot, err = filepath.Abs(websiteroot)
	if err != nil {
		return
	}
	//fmt.Println(websiteroot)
	//return

	words := vfs.MapKeys()
	for _, word := range words {
		err = CreateSymlink(word, websiteroot)
		if err != nil {
			return
		}
	}

	return
}
