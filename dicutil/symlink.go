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

func conditionalPrint(word string) {
	dontPrintEnv := []string{"TRAVIS", "GITLAB_CI"}
	for _, ev := range dontPrintEnv {
		if _, ok := os.LookupEnv(ev); ok {
			return
		}
	}
	fmt.Println(word)
}

// CreatePrefixSymlink create symbolic links for pages of Pali words to the root
// index.html of the website root directory. This is for deploying single page
// application (SPA) on GitHub Pages or GitLab Pages, which serve only static
// website content.
//
// The URL path of prefix:
//
//   /browse/{{first char of word}}/
//
// This page contains all Pali words starts with the prefix.
func CreatePrefixSymlink(prefix, root string) (err error) {
	err = os.Chdir(root)
	if err != nil {
		return
	}

	prefixIndex := filepath.Join(lib.PrefixUrlPath(prefix), "index.html")
	// remove heading /
	prefixIndex = prefixIndex[1:]
	//fmt.Println("prefix index.html path:", prefixIndex)
	err = os.Symlink("../../index.html", prefixIndex)
	if os.IsExist(err) {
		// If the symlink we want to create already exist, error will be
		// raised. Remove the existing symlink and create new symlink.
		os.Remove(prefixIndex)
		err = os.Symlink("../../index.html", prefixIndex)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}

	conditionalPrint(prefix)
	return
}

// CreateWordSymlink create symbolic links for pages of Pali words to the root
// index.html of the website root directory. This is for deploying single page
// application (SPA) on GitHub Pages or GitLab Pages, which serve only static
// website content.
//
// The URL path of word:
//
//   /browse/{{first char of word}}/{{word}}/
//
// This page contains the definition of the Pali word.
func CreateWordSymlink(word, root string) (err error) {
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
		// raised. Remove the existing symlink and create new symlink.
		os.Remove(wordIndex)
		err = os.Symlink("../../../index.html", wordIndex)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}

	conditionalPrint(word)
	return
}

// SymlinkToRootIndexHtml creates symbolic links which points all pages of the
// website to the root index.html in the root directory of the website. The
// purpose is to deploy single page application (SPA) on GitHub Pages or GitLab
// Pages, which serves only static website contents.
//
// Only one page: ``/index.html``
//
// All other webpages are symlinks to ``/index.html``
func SymlinkToRootIndexHtml(websiteroot string) (err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	websiteroot, err = filepath.Abs(websiteroot)
	if err != nil {
		return
	}
	//fmt.Println(websiteroot)
	//return

	prefixs := make(map[string]bool)
	words := vfs.MapKeys()
	for _, word := range words {
		// collect prefix of word
		prefix := lib.GetFirstCharacterOfWord(word)
		if _, ok := prefixs[prefix]; !ok {
			prefixs[prefix] = true
		}

		err = CreateWordSymlink(word, websiteroot)
		if err != nil {
			return
		}
	}

	for prefix, _ := range prefixs {
		err = CreatePrefixSymlink(prefix, websiteroot)
		if err != nil {
			return
		}
	}

	// change back to original directory to prevent causing unwanted results
	// if users call this methods multiple times in program
	return os.Chdir(wd)
}
