package main

import (
	"fmt"

	str "github.com/mutasim77/strlib"
	st2 "github.com/mutasim77/strlib/v2"
)

func main() {
	reversed := str.Reverse("Hello, 世界! How are you?")
	fmt.Println(reversed) // Output: ?uoy era woH !界世 ,olleH

	// Count symbols in a string
	count := st2.SymbolCount("Hello, 世界!")
	fmt.Println(count) // Output: 10
}
