package tpkutil

// Download Tipiṭaka xml from https://www.tipitaka.org/romn/

import (
	"encoding/xml"
	"io/ioutil"
	"path"

	"github.com/siongui/gopalilib/util"
)

type Tree struct {
	XMLName xml.Name `xml:"tree"`
	Trees   []Tree   `xml:"tree"`
	Text    string   `xml:"text,attr"`
	Src     string   `xml:"src,attr"`
}

func GetXml(srcUrl, dstPath string, overwrite bool) (t Tree, err error) {
	util.CreateDirIfNotExist(dstPath)
	err = util.CheckDownload(srcUrl, dstPath, overwrite)
	if err != nil {
		return
	}

	b, err := ioutil.ReadFile(dstPath)
	if err != nil {
		return
	}

	err = xml.Unmarshal(b, &t)
	return
}

// DownloadTipitaka downloads all Tipiṭaka XMLs from
// https://www.tipitaka.org/romn/ to dir. This method will overwrite existing
// XMLs if overwrite set to true.
func DownloadTipitaka(dir string, overwrite bool) (err error) {
	// Python version
	// https://github.com/siongui/pali/blob/master/tipitaka/setup/init1getTocs.py

	// observation:
	//  1. All meaningful node has attribute 'text'
	//  2. node with 'action' attribute is leaf
	urlPrefix := "https://www.tipitaka.org/romn/"
	rootTocXmlSrc := "tipitaka_toc.xml"

	srcUrl := urlPrefix + rootTocXmlSrc
	dstPath := path.Join(dir, rootTocXmlSrc)

	tree, err := GetXml(srcUrl, dstPath, overwrite)
	if err != nil {
		return
	}
	util.PrettyPrint(tree)

	return
}
