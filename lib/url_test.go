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
}
