package dicutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	vfs "github.com/siongui/gopaliwordvfs"
)

func TestVFS(t *testing.T) {
	files, err := ioutil.ReadDir(wordsJsonDir)
	if err != nil {
		t.Error(err)
		return
	}

	total := 0
	for i, file := range files {
		bVfs, err := vfs.ReadFile(file.Name())
		if err != nil {
			t.Error(err)
			return
		}

		bReal, err := ioutil.ReadFile(path.Join(wordsJsonDir, file.Name()))
		if err != nil {
			t.Error(err)
			return
		}

		if !bytes.Equal(bVfs, bReal) {
			t.Error(file.Name())
			return
		}

		if _, ok := os.LookupEnv("TRAVIS"); !ok {
			fmt.Println(i, file.Name(), "ok")
		}

		total++
	}

	filenames := vfs.MapKeys()
	if len(filenames) == total {
		fmt.Println("total number of json file correct")
	} else {
		t.Error("total number of json files not correct")
		return
	}
	for _, filename := range filenames {
		p := path.Join(wordsJsonDir, filename)
		if _, err := os.Stat(p); err == nil {
			if _, ok := os.LookupEnv("TRAVIS"); !ok {
				fmt.Println(p, "exist")
			}

		} else {
			t.Error(p)
			return
		}
	}
}
