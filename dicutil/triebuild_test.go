package dicutil

import (
	"testing"
)

func TestBuildSuccinctTrie(t *testing.T) {
	BuildSuccinctTrie(wordsJsonDir, trieDataPath, trieNodeCountPath, rankDirectoryDataPath)
}
