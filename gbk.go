package qqwry

import (
	"golang.org/x/text/encoding/simplifiedchinese"
)

var gbkDecoder = simplifiedchinese.GBK.NewDecoder()

func GbkString(gbk []byte) string {
	gbkDecoder.Reset()
	utf8 := make([]byte, len(gbk)*2)
	utf8len, _, _ := gbkDecoder.Transform(utf8, gbk, false)
	return string(utf8[0:utf8len])
}
