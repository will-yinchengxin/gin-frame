package utils

import (
	"strings"
"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转大驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, v := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(v))
			continue
		}
		if unicode.IsUpper(v) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(v))
	}
	return string(output)
}
