package strutils

/*
	Symbols include letters, numbers, emojis, and other characters.
	Converting the string to a slice of runes() ensures correct counting of Unicode characters,
	as each rune represents a single Unicode character, including multi-byte characters.
*/

func SymbolCount(s string) int {
	count := 0
	for range []rune(s) {
		count++
	}
	return count
}
