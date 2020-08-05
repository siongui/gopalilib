package util

import (
	"encoding/json"
	"os"
)

// SaveJsonFile saves the data to JSON file.
func SaveJsonFile(v interface{}, path string) {
	fo, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	e := json.NewEncoder(fo)
	if err := e.Encode(v); err != nil {
		panic(err)
	}
}

// LoadJsonConfig loads map[string]string data from the json file. Make sure the
// given JSON file is of the struct map[string]string.
func LoadJsonConfig(fp string) (conf map[string]string, err error) {
	f, err := os.Open(fp)
	if err != nil {
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&conf)
	return
}
