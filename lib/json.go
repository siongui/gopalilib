package lib

import (
	"encoding/json"
	"io"
)

func DecodeHttpRespWord(respBody io.ReadCloser) (wi BookIdWordExps, err error) {
	dec := json.NewDecoder(respBody)
	err = dec.Decode(&wi)
	return
}

// not used
func DecodeWordJson(wordJsonStr string) (wi BookIdWordExps, err error) {
	err = json.Unmarshal([]byte(wordJsonStr), &wi)
	return
}

func PaliSettingToJsonString(setting PaliSetting) (s string, err error) {
	b, err := json.Marshal(setting)
	s = string(b)
	return
}

func JsonStringToPaliSetting(jsonStr string) (setting PaliSetting, err error) {
	err = json.Unmarshal([]byte(jsonStr), &setting)
	return
}
