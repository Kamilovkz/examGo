package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune(10)
		return
	}
	arg := os.Args[1:]
	var result string
	for _, word := range arg[0] + arg[1] {
		if !double(result, word) {
			result += string(word)
		}
	}
	for _, w := range result {
		z01.PrintRune(w)
	}
	z01.PrintRune(10)
}

func double(s string, r rune) bool {
	for _, i := range s {
		if i == r {
			return true
		}
	}
	return false
}
