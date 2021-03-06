package dicutil

// Parse the two CSV files containing definitions of words.

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/util"
)

// Given dictionary book id, return if the dictionary book is Chinese.
func isChineseDictionary(id string) bool {
	// id of Chinese Dictionary: D G Z X H W F T J M
	switch id {
	case "D", "G", "Z", "X", "H", "W", "F", "T", "J", "M":
		return true
	default:
		return false
	}
}

func isBurmeseDictionary(id string) bool {
	switch id {
	case "B", "K", "O", "R":
		return true
	default:
		return false
	}
}

func getWordPath(word, wordsJsonDir string) string {
	return wordsJsonDir + "/" + word + ".json"
}

func getBookIdWordExps(word, wordsJsonDir string) lib.BookIdWordExps {
	f, err := os.Open(getWordPath(word, wordsJsonDir))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wi, err := lib.DecodeHttpRespWord(f)
	if err != nil {
		panic(err)
	}
	return wi
}

func appendToBookIdWordExps(wi lib.BookIdWordExps, bookId, explanation string) lib.BookIdWordExps {
	if isChineseDictionary(bookId) {
		// convert simplified chinese to traditional chinese
		wi = append(wi, [2]string{bookId, util.S2T(explanation)})
		return wi
	}

	if isBurmeseDictionary(bookId) {
		// convert Zawgyi to unicode
		wi = append(wi, [2]string{bookId, util.Zg2uni(explanation)})
		return wi
	}

	wi = append(wi, [2]string{bookId, explanation})
	return wi
}

// The format of record in dict_words_1.csv and dict_words_2.csv:
//
// row = [cell1, cell2, cell3, cell4, cell5, cell6, cell7], each row represent
// represents the explanation of a pali word in one dictionary.
//
// cell1: number of the row.
//
// cell2: the same as cell1 of dict-books.csv.
// "C" means chinese dictionary, "E" means non-chinese dictionary.
//
// cell3: the same as cell2 of dict-books.csv.
// id of the dictionary. Each dictionary has a unique value.
//
// cell4: fuzzy spelling of the pali word
//
// cell5 and cell6: the pali word. The first character of the cell may be
// upper-case.
//
// cell7: the explanation of the pali word in one dictionary.
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

	util.LocalPrintln(num, word)

	// Google search: golang check if file exists
	path := getWordPath(word, wordsJsonDir)
	if _, err := os.Stat(path); err == nil {
		// append new data to existing json file
		wi := getBookIdWordExps(word, wordsJsonDir)
		wi = appendToBookIdWordExps(wi, bookId, explanation)
		util.SaveJsonFile(wi, path)
	} else {
		// create new json file
		wi := lib.BookIdWordExps{}
		wi = appendToBookIdWordExps(wi, bookId, explanation)
		util.SaveJsonFile(wi, path)
	}
}

// Parse dict_words_1.csv and dict_words_2.csv
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
	util.CreateDirIfNotExist(outputdir)
	processWordsCSV(csv1, outputdir)
	processWordsCSV(csv2, outputdir)
}
