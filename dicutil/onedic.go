package dicutil

// Get one dictionary from two CSV files containing definitions of words.

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/siongui/gopalilib/util"
)

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
func getFromCSVRecord(record []string, bId string) (result []string) {
	// number of the word, useless
	num := record[0]

	// id of the book which the word belongs to
	bookId := record[2]

	// word (The first character of the cell may be upper-case)
	// Google search: golang lowercase
	word := strings.ToLower(record[4])

	// explanation of the pali word in one dictionary
	explanation := record[6]

	if bId == bookId {
		util.LocalPrintln(num, word, explanation)
		result = append(result, word)
		result = append(result, explanation)
	}
	return
}

// Parse dict_words_1.csv and dict_words_2.csv
func getOneDicFromWordsCSV(csvPath, bookId string) (result [][]string) {
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
		r := getFromCSVRecord(record, bookId)
		if len(r) > 0 {
			result = append(result, r)
		}
	}

	return
}

func SaveCsv(data [][]string, bookId string) {
	file, err := os.Create(bookId + ".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			panic(err)
		}
	}
}

func GetOneDictionay(csv1, csv2, bookId string) {
	r1 := getOneDicFromWordsCSV(csv1, bookId)
	r2 := getOneDicFromWordsCSV(csv2, bookId)

	r := [][]string{}
	r = append(r, r1...)
	r = append(r, r2...)
	SaveCsv(r, bookId)
}
