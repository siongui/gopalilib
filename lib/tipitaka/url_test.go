package tipitaka

import (
	"testing"
)

func TestActionToPaliTextPath(t *testing.T) {
	s := ActionToPaliTextPath("cscd/vin01m.mul2.xml")
	if s != "/romn/cscd/vin01m/mul2/" {
		t.Error(s)
		return
	}
}

func TestPaliTextPathToActionMap(t *testing.T) {
	for k, v := range actionToPaliTextPathMap {
		if k2, ok := paliTextPathToActionMap[v]; !ok && k2 != k {
			t.Error(k, v, k2)
			return
		}
	}
}

func TestDeterminePageType(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if DeterminePageType("/") != RootPage {
		t.Error("/")
	}
	if DeterminePageType("/pali-dictionary/") != NoSuchPage {
		t.Error("/pali-dictionary/")
	}
	if DeterminePageType("/romn/cscd/vin01m/mul2/") != PaliTextPage {
		t.Error("/romn/cscd/vin01m/mul2/")
	}
	if DeterminePageType("/pali-dictionary/cscd/vin01m/mul2/") != NoSuchPage {
		t.Error("/pali-dictionary/cscd/vin01m/mul2/")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")
	// FIXME: make the following test pass
	//if DeterminePageType("/") != NoSuchPage {
	//	t.Error("/")
	//}
	if DeterminePageType("/pali-dictionary/") != RootPage {
		t.Error("/pali-dictionary/")
	}
	if DeterminePageType("/pali-dictionary/romn/cscd/vin01m/mul2/") != PaliTextPage {
		t.Error("/pali-dictionary/romn/cscd/vin01m/mul2/")
	}
	if DeterminePageType("/pali-dictionary/zh_TW/cscd/vin01m/mul2/") != NoSuchPage {
		t.Error("/pali-dictionary/zh_TW/cscd/vin01m/mul2/")
	}

	SetCurrentLocale("zh_TW")
	if DeterminePageType("/pali-dictionary/zh_TW/") != RootPage {
		t.Error("/pali-dictionary/zh_TW/")
	}
	// FIXME: make the following test pass
	//if DeterminePageType("/pali-dictionary/cscd/vin01m/mul2/") != NoSuchPage {
	//	t.Error("/pali-dictionary/cscd/vin01m/mul2/")
	//}
	if DeterminePageType("/pali-dictionary/zh_TW/romn/cscd/vin01m/mul2/") != PaliTextPage {
		t.Error("/pali-dictionary/zh_TW/romn/cscd/vin01m/mul2/")
	}
	if DeterminePageType("/pali-dictionary/zh_TW/abc/cscd/vin01m/mul2/") != NoSuchPage {
		t.Error("/pali-dictionary/zh_TW/abc/cscd/vin01m/mul2/")
	}
}

func TestIsValidPaliTextUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if IsValidPaliTextUrlPath("/romn/cscd/vin01m/mul2/") != true {
		t.Error("/romn/cscd/vin01m/mul2/")
	}

	if IsValidPaliTextUrlPath("/abc/cscd/vin01m/mul2/") != false {
		t.Error("/abc/cscd/vin01m/mul2/")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")
	if IsValidPaliTextUrlPath("/pali-dictionary/romn/cscd/vin01m/mul2/") != true {
		t.Error("/pali-dictionary/romn/cscd/vin01m/mul2/")
	}

	SetCurrentLocale("zh_TW")
	if IsValidPaliTextUrlPath("/pali-dictionary/zh_TW/romn/cscd/vin01m/mul2/") != true {
		t.Error("/pali-dictionary/zh_TW/romn/cscd/vin01m/mul2/")
	}
}

func TestActionToUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if ActionToUrlPath("cscd/vin01m.mul2.xml") != "/romn/cscd/vin01m/mul2/" {
		t.Error("cscd/vin01m.mul2.xml")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")
	if ActionToUrlPath("cscd/vin01m.mul2.xml") != "/pali-dictionary/romn/cscd/vin01m/mul2/" {
		t.Error("cscd/vin01m.mul2.xml")
	}

	SetCurrentLocale("zh_TW")
	if ActionToUrlPath("cscd/vin01m.mul2.xml") != "/pali-dictionary/zh_TW/romn/cscd/vin01m/mul2/" {
		t.Error("cscd/vin01m.mul2.xml")
	}
}
