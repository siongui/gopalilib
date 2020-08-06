package lib

import (
	"encoding/json"
	"io"
)

// DecodeHttpRespWord decodes BookIdWordExps struct from HTTP response body.
func DecodeHttpRespWord(respBody io.ReadCloser) (wi BookIdWordExps, err error) {
	dec := json.NewDecoder(respBody)
	err = dec.Decode(&wi)
	return
}

// PaliSettingToJsonString saves PaliSetting struct to JSON string.
func PaliSettingToJsonString(setting PaliSetting) (s string, err error) {
	b, err := json.Marshal(setting)
	s = string(b)
	return
}

// JsonStringToPaliSetting loads PaliSetting struct from JSON string.
func JsonStringToPaliSetting(jsonStr string) (setting PaliSetting, err error) {
	err = json.Unmarshal([]byte(jsonStr), &setting)
	return
}
