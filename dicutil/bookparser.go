// Package dicutil provides methods for offline processing to build Pāli
// Dictionary.
// This package is for offline processing and not used in frontend code.
package dicutil

// In this script, we will parse information about dictionaries and build the
// type "BookIdAndInfos" struct, and save the information in JSON file.
//
// References:
// https://www.google.com/search?q=golang+read+csv

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/util"
)

// Parse the record containing information of one dictionary book.
//
// The format of record in dict-books.csv:
//
// row = [cell1, cell2, cell3, cell4], each row represents one dictionary.
//
// cell1 (b_lang): "C" means chinese dictionary,
// "E" means non-chinese dictionary.
//
// cell2 (b_num): id of the dictionary. Each dictionary has a unique value.
//
// cell3 (b_dict): name of the dictionary.
//
// cell4 (b_dictauthor): name and author of the dictionary.
func parseRecord(record []string) (id string, dict lib.BookInfo) {
	// language of the dictionary,
	// "C" means Chinese and Japanese dictionary,
	// "E" means non-Chinese dictionary.
	lang := record[0]
	// id of the dictionary. Each dictionary has a unique value.
	id = record[1]
	// short name of the dictionary.
	name := record[2]
	// name and author of the dictionary.
	author := record[3]

	switch lang {
	case "C":
		// Chinese and Japanese dictionaries
		switch id {
		case "A":
			// Japanese dictionary
			dict.Lang = "ja"
			dict.Separator = " -"
			dict.Name = "《パーリ語辞典》"
			dict.Author = "増補改訂パーリ語辞典  水野弘元著"
		case "S":
			// Japanese dictionary
			dict.Lang = "ja"
			dict.Separator = " -"
			dict.Name = "《パーリ語辞典》"
			dict.Author = "パーリ語辞典  水野弘元著"
		default:
			// Chinese dictionary
			dict.Lang = "zh"

			switch id {
			case "D":
				dict.Separator = "~"
			case "H":
				dict.Separator = " -"
			case "T":
				dict.Separator = " -"
			default:
				dict.Separator = "。"
			}

			dict.Name = util.S2T(name)
			dict.Author = util.S2T(author)
		}
	case "E":
		// English, Vietnam, Myanmar dictionaries
		switch id {
		case "U", "Q", "E":
			// Vietnamese dictionary
			dict.Lang = "vi"
			// FIXME: is "。" correct separator?
			dict.Separator = "。"
		case "B", "K", "O", "R":
			// Burmese(Myanmar) dictionary
			dict.Lang = "my"
			// FIXME: is "。" correct separator?
			dict.Separator = "。"
			dict.Name = util.Zg2uni(name)
			dict.Author = util.Zg2uni(author)
			return
		default:
			// English dictionary
			dict.Lang = "en"
			switch id {
			case "N":
				dict.Separator = "<br>"
			case "C":
				dict.Separator = "<br>"
			case "P":
				dict.Separator = "<i>"
			default:
				dict.Separator = "。"
			}
		}
		dict.Name = name
		dict.Author = author
	default:
		panic("wrong lang")
	}
	return
}

// Input is CSV file containing information of dictionary books of all
// languages. Output is parsed result and save in JSON file.
func ParseDictionayBookCSV(inBookCsv, outJson string) {
	// open csv file
	fcsv, err := os.Open(inBookCsv)
	if err != nil {
		panic(err)
	}
	defer fcsv.Close()

	// read csv
	di := lib.BookIdAndInfos{}
	r := csv.NewReader(fcsv)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if record[0] == "b_lang" {
			continue
		}
		id, dict := parseRecord(record)
		di[id] = dict
	}

	// save parsed result in JSON file
	util.CreateDirIfNotExist(outJson)
	util.SaveJsonFile(di, outJson)

	util.PrettyPrint(di)
}
