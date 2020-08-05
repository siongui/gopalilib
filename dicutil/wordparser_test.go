package dicutil

import (
	"flag"
	"testing"
)

var WordCSV1 = flag.String("WordCSV1", ".", "csv file of pali words #1")
var WordCSV2 = flag.String("WordCSV2", ".", "csv file of pali words #2")
var wordsJsonDir = flag.String("wordsJsonDir", ".", "output dir of json files of pali words")

func TestParseDictionayWordCSV(t *testing.T) {
	ParseDictionayWordCSV(*WordCSV1, *WordCSV2, *wordsJsonDir)
}
