package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	sentence := os.Args[len(os.Args)-1]
	for _, word := range sentence {
		z01.PrintRune(word)
	}
	z01.PrintRune(10)
}
