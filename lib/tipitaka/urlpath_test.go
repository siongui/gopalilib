package tipitaka

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/siongui/gopalilib/lib"
)

func TestTrimTreeText(t *testing.T) {
	s := TrimTreeText("1. pārājikakaṇḍaṃ")
	if s != "pārājikakaṇḍaṃ" {
		t.Error(s)
		return
	}

	s = TrimTreeText("12. Sattasatikakkhandhakaṃ")
	if s != "sattasatikakkhandhakaṃ" {
		t.Error(s)
		return
	}

	s = TrimTreeText("Sāratthadīpanī-ṭīkā-3")
	if s != "sāratthadīpanī" {
		t.Error(s)
		return
	}

	s = TrimTreeText("5. Pāṭidesanīyakaṇḍaṃ (bhikkhunīvibhaṅgavaṇṇanā)")
	if s != "pāṭidesanīyakaṇḍaṃ (bhikkhunīvibhaṅgavaṇṇanā)" {
		t.Error(s)
		return
	}

	s = TrimTreeText("Dhātukathāpakaraṇa-mūlaṭīkā")
	if s != "dhātukathāpakaraṇa" {
		t.Error(s)
		return
	}

	s = TrimTreeText("Kathāvatthupakaraṇa-anuṭīkā")
	if s != "kathāvatthupakaraṇa-anuṭīkā" {
		t.Error(s)
		return
	}

	s = TrimTreeText("Dīghanikāya (ṭīkā)")
	if s != "dīgha" {
		t.Error(s)
		return
	}

	s = TrimTreeText("Pārājikakaṇḍa-aṭṭhakathā")
	if s != "pārājika" {
		t.Error(s)
		return
	}

	s = TrimTreeText("13-18. anulomapaccanīyapaṭṭhānavaṇṇanā")
	if s != "anulomapaccanīyapaṭṭhānavaṇṇanā" {
		t.Error(s)
		return
	}

	s = TrimTreeText("Dīgha nikāya (aṭṭhakathā)")
	if s != "dīgha" {
		t.Error(s)
		return
	}

	s = TrimTreeText("bhikkhupātimokkhapāḷi")
	if s != "bhikkhupātimokkha" {
		t.Error(s)
		return
	}

	s = TrimTreeText("Abhinīhāra kathā")
	if s != "abhinīhāra kathā" {
		t.Error(s)
		return
	}

	// TODO: add more test cases
	s = TrimTreeText("")
	if s != "" {
		t.Error(s)
		return
	}
}

func TestTrimTreeText2(t *testing.T) {
	s := TrimTreeText2("Abhinīhāra kathā")
	if s != "abhinīhāra" {
		t.Error(s)
		return
	}
}

func TestTraverseTreeAndSetSubpathProperty(t *testing.T) {
	b, err := ioutil.ReadFile("tpktoc.json")
	if err != nil {
		t.Error(err)
		return
	}

	tree := lib.Tree{}
	err = json.Unmarshal(b, &tree)
	if err != nil {
		t.Error(err)
		return
	}

	TraverseTreeAndSetSubpathProperty(tree)
}
