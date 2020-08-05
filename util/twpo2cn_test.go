package util

import (
	"testing"
)

func TestTwPoToCN(t *testing.T) {
	localeDir := "../locale/"
	twpopath := localeDir + "zh_TW/LC_MESSAGES/messages.po"
	cnpopath := localeDir + "zh_CN/LC_MESSAGES/messages.po"

	TwPoToCN(twpopath, cnpopath)
}
