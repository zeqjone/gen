package utils

import (
	"regexp"
	"strings"
)

var (
	regCamel2Snake  = regexp.MustCompile(`[A-Z]`)
	regSnake2Camel  = regexp.MustCompile(`_[a-z]`)
	funcCamel2Snake = func(s string) string {
		return "_" + strings.ToLower(s)
	}
	funcSnake2Camel = func(s string) string {
		bt := []byte(s)
		return strings.ToUpper(string(bt[1]))
	}
)

// Snake2Camel 蛇形命名换成驼峰命名
func Snake2Camel(s string) string {
	return regSnake2Camel.ReplaceAllStringFunc(s, funcSnake2Camel)
}

// Snake2Camel 蛇形命名换成驼峰命名
func Snake2Pascal(s string) string {
	scamel := regSnake2Camel.ReplaceAllStringFunc(s, funcSnake2Camel)
	return strings.ToUpper(scamel[:1]) + scamel[1:]
}

// Camel2Snake 驼峰命名换成蛇形命名
func Camel2Snake(s string) string {
	return regCamel2Snake.ReplaceAllStringFunc(s, funcCamel2Snake)
}
