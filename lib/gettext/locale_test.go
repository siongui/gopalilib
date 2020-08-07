package gettext

import (
	"testing"
)

func TestDetermineLocaleByNavigatorLanguages(t *testing.T) {
	supportLocales := []string{"en_US", "vi_VN", "zh_TW", "fr_FR"}
	if DetermineLocaleByNavigatorLanguages("en-US,en,zh-TW,zh", supportLocales) != "en_US" {
		t.Error("en-US,en,zh-TW,zh")
	}
	if DetermineLocaleByNavigatorLanguages("zh", supportLocales) != "zh_TW" {
		t.Error("zh")
	}
	if DetermineLocaleByNavigatorLanguages("vi-VN,vi", supportLocales) != "vi_VN" {
		t.Error("vi-VN,vi")
	}
	if DetermineLocaleByNavigatorLanguages("de-DE,de", supportLocales) != "en_US" {
		t.Error("de-DE,de")
	}
	if DetermineLocaleByNavigatorLanguages("zh-TW,zh,en-US,en", supportLocales) != "zh_TW" {
		t.Error("zh-TW,zh,en-US,en")
	}
	if DetermineLocaleByNavigatorLanguages("zh-CN,zh,en-US,en", supportLocales) != "zh_TW" {
		t.Error("zh-CN,zh,en-US,en")
	}
	if DetermineLocaleByNavigatorLanguages("ja-JP,ja", supportLocales) != "en_US" {
		t.Error("ja-JP,ja")
	}
}
