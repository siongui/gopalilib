package dicutil

import (
	"path"
)

const DataRepoDir = "../"

var LocaleDir = path.Join("../locale")
var WordCSV1 = path.Join(DataRepoDir, "data/dictionary/dict_words_1.csv")
var WordCSV2 = path.Join(DataRepoDir, "data/dictionary/dict_words_2.csv")

const wordsJsonDir = "/tmp/paliwords/"

// for trie build
const trieDataPath = "/tmp/strie.txt"
const trieNodeCountPath = "/tmp/strie_node_count.txt"
const rankDirectoryDataPath = "/tmp/rd.txt"
