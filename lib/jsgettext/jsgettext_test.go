package jsgettext

import (
	"testing"
)

func TestGettext(t *testing.T) {
	s := Gettext("zh_TW", "Pāli Dictionary")
	if s != "巴利字典" {
		t.Error(s)
		return
	}
}
