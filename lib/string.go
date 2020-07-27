package lib

import (
	"unicode/utf8"
)

// RemoveLastChar removes the last character of input string.
func RemoveLastChar(str string) string {
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		return str[:len(str)-size]
	}
	return str
}
