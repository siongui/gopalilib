package lib

import (
	"testing"
)

func TestSetSiteUrl(t *testing.T) {
	SetSiteUrl("https://dictionary.sutta.org/")
	if rootPath != "" {
		t.Error(rootPath)
	}

	SetSiteUrl("https://dictionary.sutta.org")
	if rootPath != "" {
		t.Error(rootPath)
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")
	if rootPath != "/pali-dictionary" {
		t.Error(rootPath)
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary")
	if rootPath != "/pali-dictionary" {
		t.Error(rootPath)
	}
}

func TestSetCurrentLocale(t *testing.T) {
	SetCurrentLocale("zh_TW")
	if currentLocale != "zh_TW" {
		t.Error(currentLocale)
	}
}

func TestDeterminePageType(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if DeterminePageType("/") != RootPage {
		t.Error("error root page type")
	}

	if DeterminePageType("/about/") != AboutPage {
		t.Error("error about page type")
	}

	if DeterminePageType("/browse/s/sacca/") != WordPage {
		t.Error("error word page type")
	}

	if DeterminePageType("/browse/ā/āpadā/") != WordPage {
		t.Error("error word page type")
	}

	if DeterminePageType("/browse/s/āpadā/") == WordPage {
		t.Error("error word page type")
	}

	if DeterminePageType("/browse/s/sacca") != NoSuchPage {
		t.Error("error no such page")
	}

	if DeterminePageType("/about/something") != NoSuchPage {
		t.Error("error no such page")
	}

	if DeterminePageType("/browse/s/") != PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/āa/") == PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/ā/") != PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/%E1%B8%8D/") != PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/%E1%B8%8/") == PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/%E1%B8%8D/%E1%B8%8Dibhi/") != WordPage {
		t.Error("error word page type")
	}

	if DeterminePageType("/browse/%E1%B8%8/%E1%B8%8Dibhi/") == WordPage {
		t.Error("error word page type")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")

	if DeterminePageType("/pali-dictionary/") != RootPage {
		t.Error("error root page type")
	}

	if DeterminePageType("/pali-dictionary/about/") != AboutPage {
		t.Error("error about page type")
	}

	if DeterminePageType("/pali-dictionary/browse/ā/") != PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/pali-dictionary/browse/ā/āpadā/") != WordPage {
		t.Error("error prefix page type")
	}

	SetCurrentLocale("zh_TW")

	if DeterminePageType("/pali-dictionary/zh_TW/") != RootPage {
		t.Error("error root page type")
	}

	if DeterminePageType("/pali-dictionary/zh_TW/about/") != AboutPage {
		t.Error("error about page type")
	}

	if DeterminePageType("/pali-dictionary/zh_TW/browse/ā/") != PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/pali-dictionary/zh_TW/browse/ā/āpadā/") != WordPage {
		t.Error("error prefix page type")
	}
}

func TestIsValidPrefixUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if !IsValidPrefixUrlPath("/browse/ā/") {
		t.Error("/browse/ā/ should be true")
	}

	if IsValidPrefixUrlPath("/browse/āa/") {
		t.Error("/browse/āa/ should be false")
	}

	if IsValidPrefixUrlPath("/about/") {
		t.Error("/about/ should be false")
	}

	if !IsValidPrefixUrlPath("/browse/%E1%B8%8D/") {
		println("/browse/%E1%B8%8D/ (%E1%B8%8D is ḍ) should be true")
		t.Error("/browse/ḍ/ should be true")
	}

	if IsValidPrefixUrlPath("/browse/%E1%B8%8/") {
		println("/browse/%E1%B8%8/ should be false")
		t.Error("should be true")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")

	if IsValidPrefixUrlPath("/pali-dictionary/") {
		t.Error("/pali-dictionary/ should be false")
	}

	if IsValidPrefixUrlPath("/pali-dictionary/about/") {
		t.Error("/pali-dictionary/about/ should be false")
	}

	if !IsValidPrefixUrlPath("/pali-dictionary/browse/ā/") {
		t.Error("/pali-dictionary/browse/ā/ should be true")
	}

	if IsValidPrefixUrlPath("/pali-dictionary/browse/ā/āpadā/") {
		t.Error("/pali-dictionary/browse/ā/āpadā/ should be false")
	}

	SetCurrentLocale("zh_TW")

	if IsValidPrefixUrlPath("/pali-dictionary/zh_TW/") {
		t.Error("/pali-dictionary/ should be false")
	}

	if IsValidPrefixUrlPath("/pali-dictionary/zh_TW/about/") {
		t.Error("/pali-dictionary/zh_TW/about/ should be false")
	}

	if !IsValidPrefixUrlPath("/pali-dictionary/zh_TW/browse/ā/") {
		t.Error("/pali-dictionary/zh_TW/browse/ā/ should be true")
	}

	if IsValidPrefixUrlPath("/pali-dictionary/zh_TW/browse/ā/āpadā/") {
		t.Error("/pali-dictionary/zh_TW/browse/ā/āpadā/ should be false")
	}
}

func TestIsValidWordUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if !IsValidWordUrlPath("/browse/ā/āpadā/") {
		t.Error("/browse/ā/āpadā/ should be true")
	}

	if IsValidWordUrlPath("/browse/ā/āpadā") {
		t.Error("/browse/ā/āpadā should be false")
	}

	if IsValidWordUrlPath("/about/") {
		t.Error("/about/ should be false")
	}

	if !IsValidWordUrlPath("/browse/%E1%B8%8D/%E1%B8%8Dibhi/") {
		t.Error("/browse/ḍ/ḍibhi/ should be true")
	}

	if IsValidWordUrlPath("/browse/%E1%B8%8/%E1%B8%8Dibhi/") {
		t.Error("should be false")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")

	if IsValidWordUrlPath("/pali-dictionary/") {
		t.Error("/pali-dictionary/ should be false")
	}

	if IsValidWordUrlPath("/pali-dictionary/about/") {
		t.Error("/pali-dictionary/about/ should be false")
	}

	if IsValidWordUrlPath("/pali-dictionary/browse/ā/") {
		t.Error("/pali-dictionary/browse/ā/ should be false")
	}

	if !IsValidWordUrlPath("/pali-dictionary/browse/ā/āpadā/") {
		t.Error("/pali-dictionary/browse/ā/āpadā/ should be true")
	}

	SetCurrentLocale("zh_TW")

	if IsValidWordUrlPath("/pali-dictionary/zh_TW/") {
		t.Error("/pali-dictionary/ should be false")
	}

	if IsValidWordUrlPath("/pali-dictionary/zh_TW/about/") {
		t.Error("/pali-dictionary/zh_TW/about/ should be false")
	}

	if IsValidWordUrlPath("/pali-dictionary/zh_TW/browse/ā/") {
		t.Error("/pali-dictionary/zh_TW/browse/ā/ should be false")
	}

	if !IsValidWordUrlPath("/pali-dictionary/zh_TW/browse/ā/āpadā/") {
		t.Error("/pali-dictionary/zh_TW/browse/ā/āpadā/ should be true")
	}
}

func TestGetPrefixFromUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if GetPrefixFromUrlPath("/browse/s/") != "s" {
		t.Error("/browse/s/ should return s")
	}

	if GetPrefixFromUrlPath("/browse/āā/") != "" {
		t.Error(`/browse/āā/ should return ""`)
	}

	if GetPrefixFromUrlPath("/browse/ā/") != "ā" {
		t.Error(`/browse/ā/ should return "ā"`)
	}

	if GetPrefixFromUrlPath("/about/") != "" {
		t.Error(`/about/ should return ""`)
	}

	if GetPrefixFromUrlPath("/browse/%E1%B8%8D/") != "ḍ" {
		println("/browse/%E1%B8%8D/ should return ḍ")
		t.Error("should return ḍ")
	}

	if GetPrefixFromUrlPath("/browse/%E1%B8%8/") != "" {
		println(`/browse/%E1%B8%8/ should return ""`)
		t.Error(`should return ""`)
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")

	if GetPrefixFromUrlPath("/pali-dictionary/browse/s/") != "s" {
		t.Error("/pali-dictionary/browse/s/ should return s")
	}

	if GetPrefixFromUrlPath("/pali-dictionary/browse/ā/") != "ā" {
		t.Error(`/pali-dictionary/browse/ā/ should return "ā"`)
	}

	SetCurrentLocale("zh_TW")

	if GetPrefixFromUrlPath("/pali-dictionary/zh_TW/browse/s/") != "s" {
		t.Error("/pali-dictionary/zh_TW/browse/s/ should return s")
	}

	if GetPrefixFromUrlPath("/pali-dictionary/zh_TW/browse/ā/") != "ā" {
		t.Error(`/pali-dictionary/zh_TW/browse/ā/ should return "ā"`)
	}
}

func TestGetWordFromUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if GetWordFromUrlPath("/browse/s/sacca/") != "sacca" {
		t.Error("/browse/s/sacca/ should return sacca")
	}

	if GetWordFromUrlPath("/browse/s/āpadā/") != "" {
		t.Error(`/browse/s/āpadā/ should return ""`)
	}

	if GetWordFromUrlPath("/about/") != "" {
		t.Error(`/about/ should return ""`)
	}

	if GetWordFromUrlPath("/browse/%E1%B8%8D/%E1%B8%8Dibhi/") != "ḍibhi" {
		println("/browse/%E1%B8%8D/%E1%B8%8Dibhi/ should return ḍibhi")
		t.Error("should return ḍibhi")
	}

	if GetWordFromUrlPath("/browse/%E1%B8%8/%E1%B8%8Dibhi/") != "" {
		println(`/browse/%E1%B8%8/%E1%B8%8Dibhi/ should ""`)
		t.Error(`should return ""`)
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")

	if GetWordFromUrlPath("/pali-dictionary/browse/s/sacca/") != "sacca" {
		t.Error("/pali-dictionary/browse/s/sacca/ should return sacca")
	}

	if GetWordFromUrlPath("/pali-dictionary/browse/%E1%B8%8D/%E1%B8%8Dibhi/") != "ḍibhi" {
		println("/pali-dictionary/browse/%E1%B8%8D/%E1%B8%8Dibhi/ should return ḍibhi")
		t.Error("should return ḍibhi")
	}

	SetCurrentLocale("zh_TW")

	if GetWordFromUrlPath("/pali-dictionary/zh_TW/browse/s/sacca/") != "sacca" {
		t.Error("/pali-dictionary/zh_TW/browse/s/sacca/ should return sacca")
	}

	if GetWordFromUrlPath("/pali-dictionary/zh_TW/browse/%E1%B8%8D/%E1%B8%8Dibhi/") != "ḍibhi" {
		println("/pali-dictionary/zh_TW/browse/%E1%B8%8D/%E1%B8%8Dibhi/ should return ḍibhi")
		t.Error("should return ḍibhi")
	}
}

func TestWordUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if WordUrlPath("sacca") != "/browse/s/sacca/" {
		t.Error("error sacca path url")
	}

	if WordUrlPath("āpadā") != "/browse/ā/āpadā/" {
		t.Error("error āpadā path url")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")

	if WordUrlPath("sacca") != "/pali-dictionary/browse/s/sacca/" {
		t.Error("error sacca path url")
	}

	if WordUrlPath("āpadā") != "/pali-dictionary/browse/ā/āpadā/" {
		t.Error("error āpadā path url")
	}

	SetCurrentLocale("zh_TW")

	if WordUrlPath("sacca") != "/pali-dictionary/zh_TW/browse/s/sacca/" {
		t.Error(WordUrlPath("sacca"))
	}

	if WordUrlPath("āpadā") != "/pali-dictionary/zh_TW/browse/ā/āpadā/" {
		t.Error(WordUrlPath("āpadā"))
	}
}

func TestGetFirstCharacterOfWord(t *testing.T) {
	if GetFirstCharacterOfWord("sacca") != "s" {
		t.Error("sacca first char wrong")
	}

	if GetFirstCharacterOfWord("āpadā") != "ā" {
		t.Error("āpadā first char wrong")
	}

	if GetFirstCharacterOfWord("ṭakāra") != "ṭ" {
		t.Error("ṭakāra first char wrong")
	}

	if GetFirstCharacterOfWord("ṭ") != "ṭ" {
		t.Error("ṭ first char wrong")
	}

	if GetFirstCharacterOfWord("ḍ") != "ḍ" {
		t.Error("ḍ first char wrong")
	}
}

func TestPrefixUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if PrefixUrlPath("s") != "/browse/s/" {
		t.Error("prefix s url path wrong")
	}

	if PrefixUrlPath("ā") != "/browse/ā/" {
		t.Error("prefix ā url path wrong")
	}

	SetCurrentLocale("zh_TW")

	if PrefixUrlPath("s") != "/zh_TW/browse/s/" {
		t.Error("prefix s url path wrong")
	}

	if PrefixUrlPath("ā") != "/zh_TW/browse/ā/" {
		t.Error("prefix ā url path wrong")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")

	if PrefixUrlPath("s") != "/pali-dictionary/zh_TW/browse/s/" {
		t.Error("prefix s url path wrong")
	}

	if PrefixUrlPath("ā") != "/pali-dictionary/zh_TW/browse/ā/" {
		t.Error("prefix ā url path wrong")
	}
}
