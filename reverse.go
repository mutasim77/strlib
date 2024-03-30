package strutils

func Reverse(s string) string {
	// Convert the input string to a slice of runes.
	// This allows handling Unicode characters correctly.
	inputRunes := []rune(s)

	// two pointers
	start := 0
	end := len(inputRunes) - 1

	// Iterate over the slice of runes from both ends,
	// swapping characters until reaching the middle.
	for start < end {
		// Swap characters at the current positions.
		inputRunes[start], inputRunes[end] = inputRunes[end], inputRunes[start]

		// Move pointers towards the middle.
		start++
		end--
	}

	return string(inputRunes)
}
