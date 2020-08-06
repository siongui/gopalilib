package dicutil

import (
	"github.com/siongui/goef"
	"github.com/siongui/gopalilib/util"
)

// CreateMetadataGoCode create a data.go file which embed metadata directly in
// Go code.
func CreateMetadataGoCode(pkgName, metadataDir, outputGoFilePath string) error {
	util.CreateDirIfNotExist(outputGoFilePath)
	return goef.GenerateGoPackagePlainText(pkgName, metadataDir, outputGoFilePath)
}
