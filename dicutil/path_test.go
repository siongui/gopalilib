package dicutil

import (
	"path"
)

var PaligoDir = "../../paligo/"
var LocaleDir = path.Join(PaligoDir, "locale")
var TemplateDir = path.Join(PaligoDir, "theme/template")
var BookCSV = path.Join(PaligoDir, "data/dictionary/dict-books.csv")
var WordCSV1 = path.Join(PaligoDir, "data/dictionary/dict_words_1.csv")
var WordCSV2 = path.Join(PaligoDir, "data/dictionary/dict_words_2.csv")
var outBookJSON = "books.json"
