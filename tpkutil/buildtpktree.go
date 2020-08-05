package tpkutil

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/util"
)

// ReadXml decodes UTF-16 encoded XML file of Pāli Tipiṭaka and return lib.Tree
// struct. This method is used by BuildTipitakaTree method and TraverseTree
// method.
func ReadXml(filePath string) (t lib.Tree, err error) {
	f16, err := os.Open(filePath)
	if err != nil {
		return
	}

	err = util.DecodeUtf16XML(f16, &t)
	return
}

// TraverseTree traverses lib.Tree struct and constructs the Tipiṭaka ToC
// (Table of Content) tree. This method is used by BuildTipitakaTree method.
func TraverseTree(t *lib.Tree, dir string, layer int) (err error) {
	if t.Text == "" {
		for i, _ := range t.SubTrees {
			TraverseTree(&t.SubTrees[i], dir, layer+2)
		}
		return
	}

	if t.Src == "" {
		fmt.Printf("%sText: %s, Action: %s\n", strings.Repeat(" ", layer), t.Text, t.Action)
		for i, _ := range t.SubTrees {
			TraverseTree(&t.SubTrees[i], dir, layer+2)
		}
		return
	}

	if t.Src != "" {
		fmt.Printf("%sText: %s, Src: %s\n", strings.Repeat(" ", layer), t.Text, t.Src)
		xmlSrc := path.Join(dir, t.Src)
		newtree, err := ReadXml(xmlSrc)
		if err != nil {
			return err
		}

		if newtree.Text == "" {
			t.SubTrees = newtree.SubTrees
			for i, _ := range t.SubTrees {
				TraverseTree(&t.SubTrees[i], dir, layer+2)
			}
		}
	}
	return
}

// BuildTipitakaTree create ToC (Table of Content) tree of Tipiṭaka
func BuildTipitakaTree(dir string) (t lib.Tree, err error) {
	rootTocXmlSrc := "tipitaka_toc.xml"
	//fmt.Println(dir)
	xmlSrc := path.Join(dir, rootTocXmlSrc)
	fmt.Println(xmlSrc)

	t, err = ReadXml(xmlSrc)
	if err != nil {
		return
	}
	err = TraverseTree(&t, dir, 0)
	//util.PrettyPrint(t)
	return
}
