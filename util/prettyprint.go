package util

import (
	"encoding/json"
	"fmt"
)

// Pretty print Go variable (struct, map, array, slice, etc.).
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
