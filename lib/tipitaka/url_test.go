package tipitaka

import (
	"testing"
)

func TestActionToUrlPath(t *testing.T) {
	s := ActionToUrlPath("cscd/vin01m.mul2.xml")
	if s != "/cscd/vin01m/mul2" {
		t.Error(s)
		return
	}
}
