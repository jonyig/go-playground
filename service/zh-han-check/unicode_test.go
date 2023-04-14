package zh_han_check

import (
	"fmt"
	"testing"
)

func TestIsChineseChar(t *testing.T) {
	testCases := []struct {
		input    rune
		expected bool
	}{
		{'你', true},
		{'好', true},
		{'喜', true},
		{'欢', true},
		{'中', true},
		{'文', true},
		{'a', false},
		{'B', false},
		{'0', false},
		{'9', false},
		{' ', false},
		{'$', false},
		{'一', true},
		{'-', false},
		{'！', false},
		{'!', false},
		{'😀', false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %q", tc.input), func(t *testing.T) {
			output := IsChineseChar(tc.input)
			if output != tc.expected {
				t.Errorf("Expected: %v; Got: %v", tc.expected, output)
			}
		})
	}
}

func TestIsAllChineseText(t *testing.T) {
	tests := []struct {
		text     string
		expected bool
	}{
		{"周杰伦", true},
		{"杰伦 Chou", false},
		{"汉字", true},
		{"中文", true},
		{"123", false},
		{"abc", false},
		{"", false},
	}

	for _, tt := range tests {
		result := IsAllChineseText(tt.text)
		if result != tt.expected {
			t.Errorf("IsAllChineseText(%q): expected %t, but got %t",
				tt.text, tt.expected, result)
		}
	}
}

func TestIsHasChineseText(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"Hello, World!", false},
		{"中文测试", true},
		{"Hello, 世界！", true},
		{"123456", false},
		{"$", false},
		{"中文测试 Hello, 世界！", true},
		{"", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %q", tc.input), func(t *testing.T) {
			output := IsHasChineseText(tc.input)
			if output != tc.expected {
				t.Errorf("Expected: %v; Got: %v", tc.expected, output)
			}
		})
	}
}

func TestAddSpace(t *testing.T) {
	tests := []struct {
		firstName string
		lastName  string
		expected  string
	}{

		{"周杰伦", "陈奕迅", "陈奕迅周杰伦"},
		{"Jay", "Chou", "Jay Chou"},
		{"周杰伦", "Jay Chou", "周杰伦 Jay Chou"},
		{"周杰伦", "", "周杰伦 "},
		{"", "陈奕迅", " 陈奕迅"},
		{"", "", " "},
		{"三四", "张", "张三四"},
		{"John", "Doe", "John Doe"},
		{"三五", "張", "張三五"},
		{"Mike", "李", "Mike 李"},
		{"Mike張", "李", "Mike張 李"},
		{"Taylor", "Swift", "Taylor Swift"},
		{"Selena", "Gomez", "Selena Gomez"},
	}

	for _, tt := range tests {
		result := AddSpace(tt.firstName, tt.lastName)
		if result != tt.expected {
			t.Errorf("addSpace(%q, %q): expected %q, but got %q",
				tt.firstName, tt.lastName, tt.expected, result)
		}
	}
}

func TestGetName(t *testing.T) {
	tests := []struct {
		firstName string
		lastName  string
		expected  string
	}{
		{"親愛的", "杰伦", "杰伦"},
		{"親愛的", "三", "三"},
		{"親愛的", "Taylor Swift", " Taylor Swift"},
		{"Dear", "Taylor Swift", " Taylor Swift"},
		{"Dear", "Mike Smith", " Mike Smith"},
		{"親愛的", "Mike Smith", " Mike Smith"},
		{"親愛的", "杰伦 Chou", " 杰伦 Chou"},
		{"Dear", "杰伦 Chou", " 杰伦 Chou"},
		{"你好", "世界", "世界"},
		{"Hello", "world", " world"},
		{"你好", "world", " world"},
		{"Hello", "世界", " 世界"},
		{"", "world", " world"},
		{"你好", "", " "},
		{"", "", " "},
	}

	for _, tt := range tests {
		result := FormatText(tt.firstName, tt.lastName)
		if result != tt.expected {
			t.Errorf("getName(%q, %q): expected %q, but got %q",
				tt.firstName, tt.lastName, tt.expected, result)
		}
	}
}
