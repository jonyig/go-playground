package zh_han_check

import (
	"fmt"
	"unicode"
)

func IsChineseChar(r rune) bool {
	// 通过范围表来判断字符是否为中文字符
	return unicode.Is(unicode.Han, r)
}

func IsAllChineseText(text string) bool {
	for _, r := range text {
		if !IsChineseChar(r) {
			return false
		}
	}
	return len(text) > 0
}

func IsHasChineseText(text string) bool {
	for _, r := range text {
		if IsChineseChar(r) {
			return true
		}
	}
	return false
}

// combination name by first name and last name
// 根据名字判断是否需要加空格
func AddSpace(firstName, lastName string) string {
	if IsAllChineseText(firstName) && IsAllChineseText(lastName) {
		// 如果两个名字都含有中文字符，则不加空格
		return fmt.Sprintf("%s%s", lastName, firstName)
	} else {
		// 其他情况都加空格
		return fmt.Sprintf("%s %s", firstName, lastName)
	}
}

func FormatText(firstText, secondText string) string {
	if IsAllChineseText(firstText) && IsAllChineseText(secondText) {
		return secondText
	}
	return " " + secondText
}
