package lib

import (
	"testing"
)

func TestDeterminePageType(t *testing.T) {
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
}

func TestIsValidWordUrlPath(t *testing.T) {
	if !IsValidWordUrlPath("/browse/ā/āpadā/") {
		t.Error("/browse/ā/āpadā/ should be true")
	}

	if IsValidWordUrlPath("/browse/ā/āpadā") {
		t.Error("/browse/ā/āpadā should be false")
	}

	if IsValidWordUrlPath("/about/") {
		t.Error("/about/ should be false")
	}
}

func TestWordUrlPath(t *testing.T) {
	if WordUrlPath("sacca") != "/browse/s/sacca/" {
		t.Error("error sacca path url")
	}

	if WordUrlPath("āpadā") != "/browse/ā/āpadā/" {
		t.Error("error āpadā path url")
	}
}
