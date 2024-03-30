package test

import (
	"testing"

	str "github.com/mutasim77/strlib"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"12345", "54321"},
		{"This is a test.", ".tset a si sihT"},
		{"123 🌟", "🌟 321"},
		{"🎉🎈🎂", "🎂🎈🎉"},
		{"漢字", "字漢"},
		{"नमस्ते", "ेत्समन"},
		{"مرحبا", "ابحرم"},
		{"😊🌍", "🌍😊"},
		{"⌘ Hello, 世界! ⚡", "⚡ !界世 ,olleH ⌘"},
		{"🚀 Let's fly to the 🌕!", "!🌕 eht ot ylf s'teL 🚀"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := str.Reverse(tc.input)
			if result != tc.expected {
				t.Errorf("Reverse(%q) = %q; expected %q", tc.input, result, tc.expected)
			}
		})
	}
}
