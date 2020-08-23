// Package velthuis provide inputting Pali words using Velthuis scheme.
package velthuis

import (
	"github.com/gopherjs/gopherjs/js"
)

var paliInputElement *js.Object
var lastInput string = ""

func checkLastTwoCharacter(lastChar, newChar string) string {
	if lastChar == "A" && newChar == "A" {
		return "Ā"
	}
	if lastChar == "a" && newChar == "a" {
		return "ā"
	}
	if lastChar == "A" && newChar == "a" {
		return "ā"
	}
	if lastChar == "I" && newChar == "I" {
		return "Ī"
	}
	if lastChar == "i" && newChar == "i" {
		return "ī"
	}
	if lastChar == "I" && newChar == "i" {
		return "ī"
	}
	if lastChar == "U" && newChar == "U" {
		return "Ū"
	}
	if lastChar == "u" && newChar == "u" {
		return "ū"
	}
	if lastChar == "U" && newChar == "u" {
		return "ū"
	}
	if lastChar == "\"" && newChar == "N" {
		return "Ṅ"
	}
	if lastChar == "\"" && newChar == "n" {
		return "ṅ"
	}
	if lastChar == "." && newChar == "M" {
		return "Ṃ"
	}
	if lastChar == "." && newChar == "m" {
		return "ṃ"
	}
	if lastChar == "~" && newChar == "N" {
		return "Ñ"
	}
	if lastChar == "~" && newChar == "n" {
		return "ñ"
	}
	if lastChar == "." && newChar == "T" {
		return "Ṭ"
	}
	if lastChar == "." && newChar == "t" {
		return "ṭ"
	}
	if lastChar == "." && newChar == "D" {
		return "Ḍ"
	}
	if lastChar == "." && newChar == "d" {
		return "ḍ"
	}
	if lastChar == "." && newChar == "N" {
		return "Ṇ"
	}
	if lastChar == "." && newChar == "n" {
		return "ṇ"
	}
	if lastChar == "." && newChar == "L" {
		return "Ḷ"
	}
	if lastChar == "." && newChar == "l" {
		return "ḷ"
	}
	return ""
}

func paliIME(event *js.Object) {
	value := paliInputElement.Get("value").String()

	if lastInput != "" && value != "" {
		v := value
		if len(v) == (len(lastInput) + 1) {
			result := checkLastTwoCharacter(lastInput[len(lastInput)-1:], v[len(v)-1:])
			if result != "" {
				paliInputElement.Set("value", v[:len(v)-2]+result)
			}
		}
	}
	lastInput = paliInputElement.Get("value").String()
}

func BindPaliInputMethodToInputTextElementById(id string) {
	paliInputElement = js.Global.Get("document").Call("getElementById", id)
	paliInputElement.Call("addEventListener", "keyup", paliIME, false)
}
