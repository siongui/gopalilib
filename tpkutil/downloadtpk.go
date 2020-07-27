package tpkutil

// Download Tipiṭaka xml from https://www.tipitaka.org/romn/

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"

	"github.com/siongui/gopalilib/util"
)

type Tree struct {
	XMLName xml.Name `xml:"tree"`
	Trees   []Tree   `xml:"tree"`
	Text    string   `xml:"text,attr"`
	Src     string   `xml:"src,attr"`
	Action  string   `xml:"action,attr"`
}

func PrintTreeInfo(t Tree) {
	fmt.Printf("Text: %s, Src: %s, Action: %s. Child #: %d\n",
		t.Text, t.Src, t.Action, len(t.Trees))
}

func DownloadAndParseXml(srcUrl, dstPath string, overwrite bool) (t Tree, err error) {
	err = util.CheckDownload(srcUrl, dstPath, overwrite)
	if err != nil {
		return
	}

	f16, err := os.Open(dstPath)
	if err != nil {
		return
	}

	err = util.DecodeUtf16XML(f16, &t)
	return
}

func ParseXmlTree(xmlTree Tree, urlPrefix, dir string, overwrite bool) (err error) {
	PrintTreeInfo(xmlTree)

	if xmlTree.Src != "" {
		// not leaf node, recursive get remaining xml
		return GetAllXml(urlPrefix, xmlTree.Src, dir, overwrite)
	}
	if xmlTree.Action != "" {
		// leaf node
		srcUrl := urlPrefix + xmlTree.Action
		dstPath := path.Join(dir, xmlTree.Action)
		err = util.CheckDownload(srcUrl, dstPath, overwrite)
		return
	}

	for _, subtree := range xmlTree.Trees {
		err = ParseXmlTree(subtree, urlPrefix, dir, overwrite)
		if err != nil {
			return
		}
	}
	return
}

func GetAllXml(urlPrefix, xmlSrc, dir string, overwrite bool) (err error) {
	srcUrl := urlPrefix + xmlSrc
	dstPath := path.Join(dir, xmlSrc)

	xmlTree, err := DownloadAndParseXml(srcUrl, dstPath, overwrite)
	if err != nil {
		return
	}
	//util.PrettyPrint(xmlTree)
	return ParseXmlTree(xmlTree, urlPrefix, dir, overwrite)
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
	if err != nil {
		return
	}

	xsl := "cscd/tipitaka-latn.xsl"
	err = util.CheckDownload(urlPrefix+xsl, path.Join(dir, xsl), overwrite)
	if err != nil {
		return
	}
	css := "cscd/tipitaka-latn.css"
	err = util.CheckDownload(urlPrefix+css, path.Join(dir, css), overwrite)
	return
}
