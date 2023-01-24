package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		for _, word := range os.Args[1] {
			if word == []rune(os.Args[2])[0] {
				word = []rune(os.Args[3])[0]
			}
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
