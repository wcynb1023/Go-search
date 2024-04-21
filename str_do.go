package main

import (
	"regexp"
	"strings"
)

func clean_str(str string) string {
	// 过滤单词连上的字符
	reg := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	str = reg.ReplaceAllString(str, "")

	// 去除多余的空格
	str = strings.Join(strings.Fields(str), " ")
	return str
}
