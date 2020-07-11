// Package dicuitl provides methods for offline processing to build Pāli
// Dictionary.
package dicutil

// This file make symlinks for Pāli Dictionary SPA (Single Page Application)
// hosted on GitHub Pages.

import (
	"fmt"
	"os"
	"path"
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
func CreateSymlink(word, root string) {
	wh := path.Join(root, lib.WordUrlPath(word)+"index.html")
	util.CreateDirIfNotExist(wh)

	err := os.Chdir(root)
	if err != nil {
		panic(err)
	}

	rp := (lib.WordUrlPath(word) + "index.html")[1:]
	err = os.Symlink("../../../index.html", rp)
	if os.IsExist(err) {
		os.Remove(rp)
		err = os.Symlink("../../../index.html", rp)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}

	printWord(word)
}

// Only one page: ``/index.html``
//
// All other webpages are symlinks to ``/index.html``
func SymlinkToRootIndexHtml(websiteroot string) (err error) {
	websiteroot, err = filepath.Abs(websiteroot)
	if err != nil {
		return
	}

	words := vfs.MapKeys()
	for _, word := range words {
		CreateSymlink(word, websiteroot)
	}

	return
}
