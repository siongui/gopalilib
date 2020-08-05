package dicutil

import (
	"flag"
	"testing"
)

var BookCSV = flag.String("BookCSV", ".", "csv file of dictionary books")
var OutputBookJSON = flag.String("OutputBookJSON", ".", "output JSON file of parsed result")

func TestParseDictionayBookCSV(t *testing.T) {
	ParseDictionayBookCSV(*BookCSV, *OutputBookJSON)
}
