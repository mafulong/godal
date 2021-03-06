package utils

import (
	"regexp"
	"strings"
	"unicode"
)

//ToCamelFirstUpper
func ToCamelFirstUpper(str string) string {
	pieces := strings.Split(str, "_")
	newPieces := make([]string, 0, len(pieces))
	for _, piece := range pieces {
		newPieces = append(newPieces, strings.Title(piece))
	}
	return strings.Join(newPieces, "")
}

// ToCamelFirstLower 首字母变小写的驼峰
func ToCamelFirstLower(str string) string {
	str = ToCamelFirstUpper(str)
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// ToSnakeCase 蛇形, 下划线分割
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
