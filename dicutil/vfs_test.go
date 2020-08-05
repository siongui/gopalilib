package dicutil

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/siongui/gopalilib/util"
	vfs "github.com/siongui/gopaliwordvfs"
)

var wordsJsonDir = flag.String("wordsJsonDir", ".", "output dir of json files of pali words")

func TestVFS(t *testing.T) {
	files, err := ioutil.ReadDir(*wordsJsonDir)
	if err != nil {
		t.Error(err)
		return
	}

	total := 0
	for i, file := range files {
		wordname := strings.TrimSuffix(file.Name(), ".json")
		bVfs, err := vfs.ReadFile(wordname)
		if err != nil {
			t.Error(err)
			return
		}

		bReal, err := ioutil.ReadFile(path.Join(*wordsJsonDir, file.Name()))
		if err != nil {
			t.Error(err)
			return
		}

		if !bytes.Equal(bVfs, bReal) {
			t.Error(file.Name(), "content not equal")
			return
		}

		if !util.IsRunOnTravisCI() && !util.IsRunOnGitLabCI() {
			fmt.Println(i, file.Name(), "ok")
		}

		total++
	}

	wordnames := vfs.MapKeys()
	if len(wordnames) == total {
		fmt.Println("total number of json file correct")
	} else {
		t.Error("total number of json files not correct")
		return
	}
	for _, wordname := range wordnames {
		p := path.Join(*wordsJsonDir, wordname+".json")
		if _, err := os.Stat(p); err == nil {
			if !util.IsRunOnTravisCI() && !util.IsRunOnGitLabCI() {
				fmt.Println(p, "exist")
			}
		} else {
			t.Error(p)
			return
		}
	}
}
