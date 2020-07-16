package dicutil

import (
	"path"
)

const PaligoDir = "../../paligo/"
const DataRepoDir = "../"

var LocaleDir = path.Join(PaligoDir, "locale")
var BookCSV = path.Join(DataRepoDir, "data/dictionary/dict-books.csv")
var WordCSV1 = path.Join(DataRepoDir, "data/dictionary/dict_words_1.csv")
var WordCSV2 = path.Join(DataRepoDir, "data/dictionary/dict_words_2.csv")

const outBookJSON = "/tmp/books.json"
const wordsJsonDir = "/tmp/paliwords/"

// for trie build
const trieDataPath = "/tmp/strie.txt"
const trieNodeCountPath = "/tmp/strie_node_count.txt"
const rankDirectoryDataPath = "/tmp/rd.txt"
