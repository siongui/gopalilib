// Package gettext provides methods for translations, similar to GNU gettext
// function.
package gettext

import (
	"encoding/json"
)

// LocalesMsg is the data structrue to store translations of PO file.
type LocalesMsg map[string]MsgIdStrPairs

// MsgIdStrPairs is the data structrue to store translations of PO file.
type MsgIdStrPairs map[string]string

var msg = LocalesMsg{}

func SetupTranslationMapping(jsonBytes []byte) (err error) {
	err = json.Unmarshal(jsonBytes, &msg)
	return
}

// Gettext translates the given string to the language specified by locale.
func Gettext(locale, str string) string {
	if val, ok := msg[locale]; ok {
		if val2, ok2 := val[str]; ok2 {
			return val2
		}
	}
	return str
}
