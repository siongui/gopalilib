package dicutil

import (
	"testing"
)

func TestParseDictionayWordCSV(t *testing.T) {
	ParseDictionayWordCSV(WordCSV1, WordCSV2, "/tmp/paliwords/")
}
