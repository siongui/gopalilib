package util

import (
	"testing"
)

func TestCheckDownload(t *testing.T) {
	CheckDownload("https://www.tipitaka.org/romn/tipitaka_toc.xml", "/tmp/tipitaka_toc.xml", false)
}
