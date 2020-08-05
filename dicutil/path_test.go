package dicutil

import (
	"path"
)

const DataRepoDir = "../"

var LocaleDir = path.Join("../locale")

// for trie build
const trieDataPath = "/tmp/strie.txt"
const trieNodeCountPath = "/tmp/strie_node_count.txt"
const rankDirectoryDataPath = "/tmp/rd.txt"
