package tpkutil

// Download Tipiṭaka xml from https://www.tipitaka.org/romn/

import (
	"fmt"
	"path"

	"github.com/siongui/gopalilib/util"
)

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

	srcXml := urlPrefix + rootTocXmlSrc
	fmt.Println(srcXml)
	dstXml := path.Join(dir, rootTocXmlSrc)
	fmt.Println(dstXml)

	err = util.CheckDownload(srcXml, dstXml, overwrite)
	return
}
