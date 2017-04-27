package util

import (
	"encoding/json"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	// FIXME: this is Example, not Test
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		t.Error("error:", err)
		return
	}

	PrettyPrint(animals)
}
