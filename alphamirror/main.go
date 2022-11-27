package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(os.Args) != 2 {
		return
	}
	for _, word := range arg[0] {
		if word >= 'a' && word <= 'z' {
			word = 'a' + 'z' - word
		}
		if word >= 'A' && word <= 'Z' {
			word = 'A' + 'Z' - word
		}
		z01.PrintRune(word)
	}
	z01.PrintRune(10)
}
