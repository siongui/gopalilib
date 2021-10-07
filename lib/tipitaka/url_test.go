package tipitaka

import (
	"testing"
)

func TestActionToCanonPath(t *testing.T) {
	s := ActionToCanonPath("cscd/vin01m.mul2.xml")
	if s != "/cscd/vin01m/mul2/" {
		t.Error(s)
		return
	}
}

func TestUrlActionMap(t *testing.T) {
	for k, v := range actionUrlMap {
		if k2, ok := urlActionMap[v]; !ok && k2 != k {
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
	if DeterminePageType("/cscd/vin01m/mul2/") != CanonPage {
		t.Error("/cscd/vin01m/mul2/")
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
	if DeterminePageType("/pali-dictionary/cscd/vin01m/mul2/") != CanonPage {
		t.Error("/pali-dictionary/cscd/vin01m/mul2/")
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
	if DeterminePageType("/pali-dictionary/zh_TW/cscd/vin01m/mul2/") != CanonPage {
		t.Error("/pali-dictionary/zh_TW/cscd/vin01m/mul2/")
	}
	if DeterminePageType("/pali-dictionary/zh_TW/abc/cscd/vin01m/mul2/") != NoSuchPage {
		t.Error("/pali-dictionary/zh_TW/abc/cscd/vin01m/mul2/")
	}
}

func TestIsValidCanonUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if IsValidCanonUrlPath("/cscd/vin01m/mul2/") != true {
		t.Error("/cscd/vin01m/mul2/")
	}

	if IsValidCanonUrlPath("/abc/cscd/vin01m/mul2/") != false {
		t.Error("/abc/cscd/vin01m/mul2/")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")
	if IsValidCanonUrlPath("/pali-dictionary/cscd/vin01m/mul2/") != true {
		t.Error("/pali-dictionary/cscd/vin01m/mul2/")
	}

	SetCurrentLocale("zh_TW")
	if IsValidCanonUrlPath("/pali-dictionary/zh_TW/cscd/vin01m/mul2/") != true {
		t.Error("/pali-dictionary/zh_TW/cscd/vin01m/mul2/")
	}
}

func TestActionToUrlPath(t *testing.T) {
	SetSiteUrl("")
	SetCurrentLocale("")

	if ActionToUrlPath("cscd/vin01m.mul2.xml") != "/cscd/vin01m/mul2/" {
		t.Error("cscd/vin01m.mul2.xml")
	}

	SetSiteUrl("https://siongui.gitlab.io/pali-dictionary/")
	if ActionToUrlPath("cscd/vin01m.mul2.xml") != "/pali-dictionary/cscd/vin01m/mul2/" {
		t.Error("cscd/vin01m.mul2.xml")
	}

	SetCurrentLocale("zh_TW")
	if ActionToUrlPath("cscd/vin01m.mul2.xml") != "/pali-dictionary/zh_TW/cscd/vin01m/mul2/" {
		t.Error("cscd/vin01m.mul2.xml")
	}
}
