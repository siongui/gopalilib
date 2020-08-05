package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Download downloads the url to local path specified by the second argument. If
// the local file already exists, it will be overwritten.
func Download(url, filePath string) (err error) {
	CreateDirIfNotExist(filePath)
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

// CheckDownload is the same as Download, except the third argument specifies
// whether to overwrite local file if it already exists.
func CheckDownload(url, filePath string, overwrite bool) (err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return Download(url, filePath)
	}
	if err == nil && overwrite {
		err = Download(url, filePath)
	}
	return
}
