package main

import (
	"fmt"

	str "github.com/mutasim77/strlib"
)

func main() {
	reversed := str.Reverse("Hello, 世界! How are you?")
	fmt.Println(reversed) // Output: ?uoy era woH !界世 ,olleH
}
