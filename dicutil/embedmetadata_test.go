package dicutil

import (
	"flag"
	"testing"
)

var metadataDir = flag.String("metadataDir", ".", "metadata dir")
var outputGoFilePath = flag.String("outputGoFilePath", "lib/dicmgr/data.go", "Go file containing data")

func TestCreateMetadataGoCode(t *testing.T) {
	err := CreateMetadataGoCode("dicmgr", *metadataDir, *outputGoFilePath)
	if err != nil {
		t.Error(err)
		return
	}
}
