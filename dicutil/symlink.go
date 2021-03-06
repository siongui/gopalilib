package dicutil

// This file make symlinks for Pāli Dictionary SPA (Single Page Application)
// hosted on GitHub Pages.

import (
	"os"
	"path/filepath"

	dic "github.com/siongui/gopalilib/lib/dictionary"
	"github.com/siongui/gopalilib/util"
)

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

	prefixIndex := filepath.Join(dic.PrefixUrlPath(prefix), "index.html")
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

	util.LocalPrintln(prefix)
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
	wordIndexAbs := filepath.Join(root, dic.WordUrlPath(word), "index.html")
	util.CreateDirIfNotExist(wordIndexAbs)
	//fmt.Println("wordIndexAbs:", wordIndexAbs)

	err = os.Chdir(root)
	if err != nil {
		return
	}

	wordIndex := filepath.Join(dic.WordUrlPath(word), "index.html")
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

	util.LocalPrintln(word)
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
func SymlinkToRootIndexHtml(websiteroot string, words []string) (err error) {
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
	for _, word := range words {
		// collect prefix of word
		prefix := dic.GetFirstCharacterOfWord(word)
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
