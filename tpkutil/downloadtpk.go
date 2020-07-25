package tpkutil

// Download Tipiṭaka xml from https://www.tipitaka.org/romn/

import (
	"fmt"
	"path"
)

func DownloadTipitaka(dir string) (err error) {
	// observation:
	//  1. All meaningful node has attribute 'text'
	//  2. node with 'action' attribute is leaf
	urlPrefix := "https://www.tipitaka.org/romn/"
	rootTocXmlSrc := "tipitaka_toc.xml"

	srcXml := path.Join(urlPrefix, rootTocXmlSrc)
	fmt.Println(srcXml)
	dstXml := path.Join(dir, rootTocXmlSrc)
	fmt.Println(dstXml)

	return
}
