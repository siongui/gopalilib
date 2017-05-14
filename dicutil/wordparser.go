package dicutil

import (
	"encoding/csv"
	"github.com/siongui/gojianfan"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/util"
	"io"
	"os"
	"strings"
)

func isChineseDictionary(id string) bool {
	// id of Chinese Dictionary: D G Z X H W F T J M
	switch id {
	case "D", "G", "Z", "X", "H", "W", "F", "T", "J", "M":
		return true
	default:
		return false
	}
}

func processWord(record []string, wordsJsonDir string) {
	// number of the word, useless
	num := record[0]

	// id of the book which the word belongs to
	bookId := record[2]

	// word (The first character of the cell may be upper-case)
	// Google search: golang lowercase
	word := strings.ToLower(record[4])

	// explanation of the pali word in one dictionary
	explanation := record[6]

	println(num + " " + word)
	// Google search: golang check if file exists
	path := GetWordPath(word, wordsJsonDir)
	if _, err := os.Stat(path); err == nil {
		// append new data to existing json file
		wi := GetBookIdWordExps(word, wordsJsonDir)
		if isChineseDictionary(bookId) {
			// convert simplified chinese to traditional chinese
			wi[bookId] = gojianfan.S2T(explanation)
		} else {
			wi[bookId] = explanation
		}
		util.SaveJsonFile(wi, path)
	} else {
		// create new json file
		wi := lib.BookIdWordExps{}
		if isChineseDictionary(bookId) {
			// convert simplified chinese to traditional chinese
			wi[bookId] = gojianfan.S2T(explanation)
		} else {
			wi[bookId] = explanation
		}
		util.SaveJsonFile(wi, path)
	}
}

func processWordsCSV(csvPath, wordsJsonDir string) {
	// open csv file
	fcsv, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}
	defer fcsv.Close()

	// read csv
	r := csv.NewReader(fcsv)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		processWord(record, wordsJsonDir)
	}
}

func ParseDictionayWordCSV(csv1, csv2, outputdir string) {
	processWordsCSV(csv1, outputdir)
	processWordsCSV(csv2, outputdir)
}
