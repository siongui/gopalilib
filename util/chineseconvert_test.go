package util

import (
	"testing"
)

func TestS2T(t *testing.T) {
	if S2T("数学") != "數學" {
		t.Error(`S2T("数学") != "數學"`)
	}
}

func TestT2S(t *testing.T) {
	if T2S("數學") != "数学" {
		t.Error(`T2S("數學") != "数学"`)
	}
}
