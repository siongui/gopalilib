package util

import (
	"encoding/json"
)

// Pretty print Go variable (struct, map, array, slice, etc.).
func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}
