package lib

import (
	"testing"
)

func HttpWordJsonPath(word string) string {
	return "https://siongui.github.io/xemaauj9k5qn34x88m4h/" + word + ".json"
}

func TestHttpGetWordJson(t *testing.T) {
	wi, err := HttpGetWordJson(HttpWordJsonPath("sacca"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(wi)
}
