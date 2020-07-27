package lib

import (
	"testing"
)

func TestRemoveLastChar(t *testing.T) {
	s := RemoveLastChar("sacca")
	if s != "sacc" {
		t.Error(s)
		return
	}

	s = RemoveLastChar("īsādantā")
	if s != "īsādant" {
		t.Error(s)
		return
	}
}
