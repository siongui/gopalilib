package lib

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// HttpGetWordJson returns BookIdWordExps struct of the word, given the url/path
// of the word JSON file.
func HttpGetWordJson(wordJsonUrl string) (wi BookIdWordExps, err error) {
	resp, err := http.Get(wordJsonUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("resp.StatusCode != 200")
		return
	}

	wi, err = DecodeHttpRespWord(resp.Body)
	return
}

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
