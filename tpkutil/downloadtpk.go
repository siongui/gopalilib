package tpkutil

// Download Tipiṭaka xml from https://www.tipitaka.org/romn/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func download(url, filePath string) (err error) {
	fmt.Println("Downloading ", url, " to ", filePath)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}

func checkDownload(url, filePath string, overwrite bool) (err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return download(url, filePath)
	}
	if err == nil && overwrite {
		err = download(url, filePath)
	}
	return
}

// DownloadTipitaka downloads all Tipiṭaka XMLs from
// https://www.tipitaka.org/romn/ to dir. This method will overwrite existing
// XMLs if overwrite set to true.
func DownloadTipitaka(dir string, overwrite bool) (err error) {
	// observation:
	//  1. All meaningful node has attribute 'text'
	//  2. node with 'action' attribute is leaf
	urlPrefix := "https://www.tipitaka.org/romn/"
	rootTocXmlSrc := "tipitaka_toc.xml"

	srcXml := urlPrefix + rootTocXmlSrc
	fmt.Println(srcXml)
	dstXml := path.Join(dir, rootTocXmlSrc)
	fmt.Println(dstXml)

	err = checkDownload(srcXml, dstXml, overwrite)
	return
}
