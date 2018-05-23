package util

import (
	"testing"
)

func TestS2T(t *testing.T) {
	if S2T("数学") != "數學" {
		t.Error(`S2T("数学") != "數學"`)
	}
}
