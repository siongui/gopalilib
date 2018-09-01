// Package dicuitl provides methods for offline processing to build Pāli
// Dictionary.
package dicutil

// This file make symlinks for Pāli Dictionary SPA (Single Page Application)
// hosted on GitHub Pages.

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/util"
)

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
	println(word)
}

// Only one page: ``/index.html``
//
// All other webpages are symlinks to ``/index.html``
func SymlinkToRootIndexHtml(jsondir, websiteroot string) (err error) {
	jsondir, err = filepath.Abs(jsondir)
	if err != nil {
		return
	}
	websiteroot, err = filepath.Abs(websiteroot)
	if err != nil {
		return
	}

	files, err := ioutil.ReadDir(jsondir)
	if err != nil {
		return
	}

	for _, file := range files {
		name := strings.TrimSuffix(file.Name(), ".json")
		CreateSymlink(name, websiteroot)
	}

	return
}
