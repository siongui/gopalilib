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
	var ies = []IdExp{
		{"Z", "Hello World。This is a test string for short explanation"},
		{"Z", "Hello World.This is a test string for short explanation......................................................................................................................................................................."},
	}

	ies2 := ShortExplanation(ies, di)

	if ies2[0].Exp != "Hello World" {
		t.Error(ies2[0].Exp)
		return
	}

	if ies2[1].Exp != "Hello World.This is a test string for short explanation...................................................................................................................................................." {
		t.Error(ies2[1].Exp)
		t.Error(len(ies2[1].Exp))
		return
	}
}
