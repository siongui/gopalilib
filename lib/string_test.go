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

func TestGetFirstCharacter(t *testing.T) {
	if GetFirstCharacter("sacca") != "s" {
		t.Error("sacca first char wrong")
	}

	if GetFirstCharacter("āpadā") != "ā" {
		t.Error("āpadā first char wrong")
	}

	if GetFirstCharacter("ṭakāra") != "ṭ" {
		t.Error("ṭakāra first char wrong")
	}

	if GetFirstCharacter("ṭ") != "ṭ" {
		t.Error("ṭ first char wrong")
	}

	if GetFirstCharacter("ḍ") != "ḍ" {
		t.Error("ḍ first char wrong")
	}
}
