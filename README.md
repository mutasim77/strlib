# StrLib 🧵

StrLib is a versatile string manipulation library for Go, designed to make working with strings easy and efficient. With StrLib, you can perform various operations such as reversing strings, counting symbols, and much more.

## 🛠️ Features
- **Reverse**: Reverse strings effortlessly with the `Reverse` function.
- **SymbolCount**: Count the number of symbols (Unicode characters) in a string using the `SymbolCount` function.

## 📦 Installation
To start using StrLib in your Go projects, simply install it using `go get`:

```go
go get github.com/mutasim77/strlib
```

For a new version:
```go
go get github.com/mutasim77/strlib/v2
```

## 💡 Usage
```go
package main

import (
	"fmt"

	str "github.com/mutasim77/strlib"
	str2 "github.com/mutasim77/strlib/v2"
)

func main() {
	reversed := str.Reverse("Hello, 世界! How are you?")
	fmt.Println(reversed) // Output: ?uoy era woH !界世 ,olleH

	// Count symbols in a string
	count := str2.SymbolCount("Hello, 世界!")
	fmt.Println(count) // Output: 10
}
```

## 📝 Versioning
StrLib follows semantic versioning (`SemVer`). Currently, the main module `github.com/mutasim77/strlib` includes only the `Reverse` function. Version 2 (`v2`) introduces a new function called `SymbolCount`.

> [!NOTE]
> If you're using the `github.com/mutasim77/strlib` module, you'll have access only to the `Reverse` function. To use the `SymbolCount` function, make sure to import version `2` of the module `github.com/mutasim77/strlib/v2`.
