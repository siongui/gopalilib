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

	tmpdir, err := ioutil.TempDir("", "vfsbuild-tmp-dir")
	if err != nil {
		return
	}
	defer os.RemoveAll(tmpdir) // clean up

	for i, file := range files {
		oldpath := path.Join(wordJsonDir, file.Name())
		newpath := path.Join(tmpdir, strings.TrimSuffix(file.Name(), ".json"))
		err = os.Link(oldpath, newpath)
		if err != nil {
			return
		}
		if !util.IsRunOnTravisCI() && !util.IsRunOnGitLabCI() {
			fmt.Println(i+1, "hard link to", oldpath, "from", newpath)
		}
	}
	err = goef.GenerateGoPackagePlainTextWithMaxFileSize(pkgName, tmpdir, outputDir, 31000000)
	return
}
