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

// GetFirstCharacter returns first character of input string. For example, āpadā
// will return ā
func GetFirstCharacter(word string) string {
	result := ""
	if len(word) > 0 {
		runeValue, _ := utf8.DecodeRuneInString(word)
		result = string(runeValue)
	}
	return result
}
