package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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

func CheckDownload(url, filePath string, overwrite bool) (err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return Download(url, filePath)
	}
	if err == nil && overwrite {
		err = Download(url, filePath)
	}
	return
}
