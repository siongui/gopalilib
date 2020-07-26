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

func GetAllXml(urlPrefix, xmlSrc, dir string, overwrite bool) (err error) {
	srcUrl := urlPrefix + xmlSrc
	dstPath := path.Join(dir, xmlSrc)

	xmlTree, err := GetXml(srcUrl, dstPath, overwrite)
	if err != nil {
		return
	}

	for _, subtree := range xmlTree.Trees {
		util.PrettyPrint(subtree)
		//util.PrettyPrint(subtree.Src)
		if subtree.Src != "" {
			err = GetAllXml(urlPrefix, subtree.Src, dir, overwrite)
			if err != nil {
				return
			}
		}
	}
	//util.PrettyPrint(xmlTree)
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

	err = GetAllXml(urlPrefix, rootTocXmlSrc, dir, overwrite)
	return
}
