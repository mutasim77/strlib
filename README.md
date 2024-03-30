# StrLib 🧵

StrLib is a Go module that provides utilities for working with strings.

## 🛠️ Features
- **ReverseString:** Turn text upside down! StrLib can flip any text you give it, making it great for secret codes or just having fun.

## 📦 Installation
To use StrLib in your projects, just tell your computer to get it:

```go
go get github.com/mutasim77/strlib
```

## 💡 Usage`
```go
package main

import (
	"fmt"
	str "github.com/mutasim77/strlib"
)

func main() {
	// Reverse a string
	reversed := str.Reverse("hello")
	fmt.Println(reversed) // Output: "olleh"
}
```