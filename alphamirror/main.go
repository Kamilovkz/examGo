package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	for _, d := range os.Args[1] {
		if d >= 'a' && d <= 'z' {
			z01.PrintRune('a' + 'z' - d)
		} else if d >= 'A' && d <= 'Z' {
			z01.PrintRune('A' + 'Z' - d)
		} else {
			z01.PrintRune(d)
		}
	}
	z01.PrintRune('\n')
}
