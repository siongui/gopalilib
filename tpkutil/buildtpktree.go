package tpkutil

import (
	"fmt"
	"os"
	"path"

	"github.com/siongui/gopalilib/util"
)

func ReadXml(filePath string) (t Tree, err error) {
	f16, err := os.Open(filePath)
	if err != nil {
		return
	}

	err = util.DecodeUtf16XML(f16, &t)
	return
}

func TraverseTree(t *Tree, dir string) (err error) {
	for i, subtree := range t.SubTrees {
		if subtree.Src != "" {
			xmlSrc := path.Join(dir, subtree.Src)
			subt, err := ReadXml(xmlSrc)
			if err != nil {
				return err
			}
			fmt.Printf("Text: %s, Src: %s, Action: %s, Child #: %d\n", subt.Text, subt.Src, subt.Action, len(subt.SubTrees))
			t.SubTrees[i].SubTrees = append(t.SubTrees[i].SubTrees, subt.SubTrees...)
			for j, _ := range t.SubTrees[i].SubTrees {
				err = TraverseTree(&t.SubTrees[i].SubTrees[j], dir)
				if err != nil {
					return err
				}
			}
		}
	}
	//fmt.Println(t)
	return
}

func BuildTipitakaTree(dir string) (err error) {
	rootTocXmlSrc := "tipitaka_toc.xml"
	//fmt.Println(dir)
	xmlSrc := path.Join(dir, rootTocXmlSrc)
	fmt.Println(xmlSrc)

	t, err := ReadXml(xmlSrc)
	if err != nil {
		return
	}
	err = TraverseTree(&t, dir)
	util.PrettyPrint(t)
	return
}
