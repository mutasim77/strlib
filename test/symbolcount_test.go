package test

import (
	"testing"

	str2 "github.com/mutasim77/strlib/v2"
)

func TestSymbolCount(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"world", 5},
		{"", 0},
		{"12345", 5},
		{"This is a test.", 15},
		{"123 🌟", 5},
		{"🎉🎈🎂", 3},
		{"漢字", 2},
		{"नमस्त", 5},
		{"مرحبا", 5},
		{"😊🌍", 2},
		{"Go is 💜", 7},
		{"⌘ Hello, 世界! ⚡", 14},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := str2.SymbolCount(tc.input)
			if result != tc.expected {
				t.Errorf("SymbolCount(%q) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}
