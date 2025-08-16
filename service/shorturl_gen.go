package service

import (
	"fmt"
	"strings"
)

const (
	dict      = "abcedfghijklmnopqrstuvwxyzABCEDFGHIJKLMNOPQRSTUVWXYZ0123456789"
	urlPrefix = "http://127.0.0.0/"
)

func GenShortUrl(url string) string {
	//todo check existing
	return fmt.Sprintf("%s%s", urlPrefix, encode(GenId()))
	//todo save
}

func encode(num int64) string {
	scale := int64(len(dict))
	var ret strings.Builder
	for num > 0 {
		idx := num % scale
		ret.WriteByte(dict[idx])
		num = num / scale
	}
	return ret.String()
}
