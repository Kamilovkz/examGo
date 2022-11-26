package main

import (
	"github.com/01-edu/z01"
)

func main() {
	sentence := "Hello World!"
	for _, word := range sentence {
		z01.PrintRune(word)
	}
	z01.PrintRune(10)
}
