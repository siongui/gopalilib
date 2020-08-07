package lib

import (
	"testing"
)

func TestGetDefaultPaliSetting(t *testing.T) {
	s := GetDefaultPaliSetting()

	if s.P2en != true {
		t.Error(s.P2en)
		return
	}
}
