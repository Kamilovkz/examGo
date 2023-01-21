package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	for _, word := range os.Args[1] {
		if word >= 'a' && word <= 'z' {
			word = 'a' + 'z' - word
		} else if word >= 'A' && word <= 'Z' {
			word = 'A' + 'Z' - word
		}
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
