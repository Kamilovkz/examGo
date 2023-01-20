package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'm' || word[i] >= 'A' && word[i] <= 'M' {
			z01.PrintRune(rune(word[i]) + 13)
		} else if word[i] >= 'n' && word[i] <= 'z' || word[i] >= 'N' && word[i] <= 'Z' {
			z01.PrintRune(rune(word[i]) - 13)
		} else {
			z01.PrintRune(rune(word[i]))
		}
	}
	z01.PrintRune('\n')
}
