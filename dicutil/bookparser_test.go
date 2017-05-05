package dicutil

import (
	"testing"
)

func TestParseDictionayBookCSV(t *testing.T) {
	ParseDictionayBookCSV(BookCSV, outBookJSON)
}
