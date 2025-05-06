package helpers

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Capitalize(str string) string {
	if str == "" {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// 数值转换
var textAttr = []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}

// Zh2Number 中文数字转阿拉伯数字
func Zh2Number(text string) int {
	num := 0
	specialMap := map[string]int{"大": 1, "小": 99}

	if val, ok := specialMap[text]; ok {
		num = val
	} else {
		// 如果以"十"开头，替换为"一十"
		if strings.HasPrefix(text, "十") {
			text = strings.Replace(text, "十", "一十", 1)
		}

		textSplits := strings.Split(text, "十")
		var dec, unit string

		if len(textSplits) == 2 {
			dec, unit = textSplits[0], textSplits[1]
		} else {
			dec, unit = "", textSplits[0]
		}

		unitVal := slices.Index(textAttr, unit)
		decVal := slices.Index(textAttr, dec)
		num = decVal*10 + unitVal
	}

	return num
}

// Number2Zh 阿拉伯数字转中文数字
func Number2Zh(num any) string {
	text := ""
	specialMap := map[int]string{1: "大", 99: "小"}

	// 处理字符串类型的数字
	var (
		numInt int
		err    error
	)
	switch v := num.(type) {
	case string:
		numInt, err = strconv.Atoi(v)
		if err != nil {
			return ""
		}
	case int:
		numInt = v
	default:
		return ""
	}

	if val, ok := specialMap[numInt]; ok {
		text = val
	} else {
		dec := numInt / 10
		unit := numInt % 10

		if dec > 0 {
			text = textAttr[dec] + "十"
			if dec == 1 {
				text = "十"
			}
		}

		text += textAttr[unit]
	}

	return text
}

// RemoveDuplicates removes duplicate elements from a slice.
func RemoveDuplicates[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	result := []T{}

	for _, item := range slice {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// ReplaceAllStringAndSpace replaces all occurrences of the pattern in the source string with the replacement string and removes spaces.
func ReplaceAllStringAndSpace(src string, pattern string, repl string) string {
	replaced := regexp.MustCompile(pattern).ReplaceAllString(src, repl)
	replaced = strings.ReplaceAll(replaced, " ", "")

	return replaced
}

// ReplaceAllStringAndSpace2 replaces all occurrences of the pattern in the source string with the replacement string and removes spaces.
func ReplaceAllStringAndSpace2(src string, pattern *regexp.Regexp, repl string) string {
	replaced := pattern.ReplaceAllString(src, repl)
	replaced = strings.ReplaceAll(replaced, " ", "")

	return replaced
}
