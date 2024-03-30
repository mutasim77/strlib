package strutils_test

import (
	"fmt"

	str "github.com/mutasim77/strlib"
)

func ExampleReverse() {
	reversed := str.Reverse("Hello, 世界!")
	fmt.Println(reversed)
	// Output: !界世 ,olleH
}
