package util

import (
	"regexp"
	"strings"
)

// TrimHTML 删除html
func TrimHTML(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, ";")
	return strings.TrimSpace(src)
}

// CutHTML 剪切预览
func CutHTML(html string, length int) string {
	if len([]rune(TrimHTML(html))) < length {
		return string([]rune(TrimHTML(html)))
	}
	return string([]rune(TrimHTML(html))[:length])
}
