package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	numbers := os.Args[1:]
	var length int
	for i := 0; i < len(numbers); i++ {
		length += 1
	}
	z01.PrintRune(rune(length + 48))
	z01.PrintRune(10)
}
