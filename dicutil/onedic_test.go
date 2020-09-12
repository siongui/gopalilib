package dicutil

import (
	"flag"
	"testing"
)

var WordCSV1 = flag.String("WordCSV1", ".", "csv file of pali words #1")
var WordCSV2 = flag.String("WordCSV2", ".", "csv file of pali words #2")

func TestGetOneDictionay(t *testing.T) {
	GetOneDictionay(*WordCSV1, *WordCSV2, "I")
}
