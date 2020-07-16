package dicutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/siongui/goef"
	"github.com/siongui/gopalilib/util"
)

func BuildVFS(pkgName, wordJsonDir, outputDir string) (err error) {
	util.CreateDirIfNotExist(outputDir)
	files, err := ioutil.ReadDir(wordJsonDir)
	if err != nil {
		return
	}

	i := 0
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			oldpath := path.Join(wordJsonDir, file.Name())
			newpath := path.Join(wordJsonDir, file.Name()[0:len(file.Name())-5])
			err = os.Rename(oldpath, newpath)
			if err != nil {
				return
			}
			if !util.IsRunOnTravisCI() {
				fmt.Println(i, "convert", oldpath, "to", newpath)
			}
		} else {
			if !util.IsRunOnTravisCI() {
				fmt.Println(i, "unchanged", file.Name())
			}
		}
		i++
	}
	err = goef.GenerateGoPackagePlainTextWithMaxFileSize(pkgName, wordJsonDir, outputDir, 31000000)
	return
}
