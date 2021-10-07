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
