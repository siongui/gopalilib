package lib

import (
	"testing"
)

func TestShortExplanation(t *testing.T) {
	// https://www.google.com/search?q=golang+map+literal
	var di = BookIdAndInfos{
		"Z": {
			Lang:      "zh",
			Separator: "。",
			Name:      "《巴漢佛學辭匯》",
			Author:    "巴利文-漢文佛學名相辭匯 翻譯：張文明",
		},
	}

	// https://www.google.com/search?q=golang+two+dimensional+literal
	// https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go
	var biwes = [][2]string{
		{"Z", "Hello World。This is a test string for short explanation"},
		{"Z", "Hello World.This is a test string for short explanation......................................................................................................................................................................."},
	}

	biwes2 := ShortExplanation(biwes, di)

	if biwes2[0][1] != "Hello World" {
		t.Error(biwes2[0][1])
		return
	}

	if biwes2[1][1] != "Hello World.This is a test string for short explanation...................................................................................................................................................." {
		t.Error(biwes2[1][1])
		t.Error(len(biwes2[1][1]))
		return
	}
}
