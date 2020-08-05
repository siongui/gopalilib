package util

import (
	"github.com/siongui/gojianfan"
)

// Pure Go implementation of OpenCC
// https://pylist.com/t/1574949866
// https://github.com/liuzl/gocc
// https://github.com/sgoby/opencc

// S2T converts Simplified Chinese to Traditional Chinese.
func S2T(s string) string {
	return gojianfan.S2T(s)
}

// T2S converts traditional Chinese to Simplified Chinese.
func T2S(s string) string {
	return gojianfan.T2S(s)
}
