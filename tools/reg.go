package tools

import (
	"regexp"
)

// @param     compile_string        string         正则表达式
// @param     text        string         文本
// @return    valid        bool         文本是否匹配
func StringValid(compile_string string, text string) bool {
	compile, err := regexp.Compile(compile_string)
	if err != nil {
		return false
	}
	return compile.MatchString(text)
}
