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
		if word >= 'a' && word <= 'n' || word >= 'A' && word <= 'N' {
			z01.PrintRune(word + 13)
		} else if word >= 'm' && word <= 'z' || word >= 'M' && word <= 'Z' {
			z01.PrintRune(word - 13)
		} else {
			z01.PrintRune(word)
		}
	}
	z01.PrintRune('\n')
}
