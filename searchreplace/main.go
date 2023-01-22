package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		for _, d := range os.Args[1] {
			if d == []rune(os.Args[2])[0] {
				d = []rune(os.Args[3])[0]
			}
			z01.PrintRune(d)
		}
		z01.PrintRune('\n')
	}
}
