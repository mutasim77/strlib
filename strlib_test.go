package strutils_test

import (
	"fmt"

	str "github.com/mutasim77/strlib"
	str2 "github.com/mutasim77/strlib/v2"
)

func ExampleReverse() {
	reversed := str.Reverse("Hello, 世界!")
	fmt.Println(reversed)
	// Output: !界世 ,olleH
}

func ExampleSymbolCount() {
	count := str2.SymbolCount("Hello, 世界!")
	fmt.Println(count)
	// Output: 10
}
