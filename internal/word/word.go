package word

import (
	"strings"
	"unicode"
)

// 转换成大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 转换成小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转成大写驼峰
func UnderscoreToUpperCamelCase(s string) string {

	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转成小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelCaseToUnderscore(s string) string {
	var output []rune

	for index, value := range s {
		if index == 0 {
			output = append(output, unicode.ToLower(value))
			continue
		}
		if unicode.IsUpper(value) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(value))
	}
	return string(output)
}
