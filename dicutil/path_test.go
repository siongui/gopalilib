package dicutil

import (
	"path"
)

var PaligoDir = "../../paligo/"
var LocaleDir = path.Join(PaligoDir, "locale")
var TemplateDir = path.Join(PaligoDir, "theme/template")
var BookCSV = path.Join(PaligoDir, "data/dictionary/dict-books.csv")
var outBookJSON = "books.json"
