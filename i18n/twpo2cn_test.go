package i18n

import (
	"testing"
)

func TestTwPoToCN(t *testing.T) {
	localeDir := "../../paligo/locale/"
	twpopath := localeDir + "zh_TW/LC_MESSAGES/messages.po"
	cnpopath := localeDir + "zh_CN/LC_MESSAGES/messages.po"

	TwPoToCN(twpopath, cnpopath)
}
