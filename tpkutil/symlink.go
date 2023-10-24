package tpkutil

// This file make symlinks for Pāli Tipiṭaka SPA (Single Page Application)
// hosted on GitHub Pages.

import (
	"os"
	"path/filepath"

	tpk "github.com/siongui/gopalilib/lib/tipitaka"
	"github.com/siongui/gopalilib/util"
)

// CreatePageSymlink create symbolic links for Pali pages to the root index.html
// of the website root directory. This is for deploying single page application
// (SPA) on GitHub Pages or GitLab Pages, which serve only static website
// content.
//
// The URL path of a Pali page looks like:
//
//   /romn/cscd/vin01m/mul0/
//
// The page contains the content of the Pali texts.
func CreatePageSymlink(pagePath, root string) (err error) {
	// create dir of the page
	pageIndexAbs := filepath.Join(root, pagePath, "index.html")
	util.CreateDirIfNotExist(pageIndexAbs)
	//fmt.Println("pageIndexAbs:", pageIndexAbs)

	err = os.Chdir(root)
	if err != nil {
		return
	}

	pageIndex := filepath.Join(pagePath, "index.html")
	// remove heading /
	pageIndex = pageIndex[1:]
	//fmt.Println("page index.html path:", pageIndex)
	err = os.Symlink("../../../../index.html", pageIndex)
	if os.IsExist(err) {
		// If the symlink we want to create already exist, error will be
		// raised. Remove the existing symlink and create new symlink.
		os.Remove(pageIndex)
		err = os.Symlink("../../../../index.html", pageIndex)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}

	util.LocalPrintln(pagePath)
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
func SymlinkToRootIndexHtml(websiteroot string, script string) (err error) {
	// FIXME TODO: script param is not respected right now. return only romn

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

	paths := tpk.GetAllCanonPath("romn")
	for _, p := range paths {
		err = CreatePageSymlink(p, websiteroot)
		if err != nil {
			return
		}
	}

	// change back to original directory to prevent causing unwanted results
	// if users call this methods multiple times in program
	return os.Chdir(wd)
}
