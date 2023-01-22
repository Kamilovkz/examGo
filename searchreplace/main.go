package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	for _, w1 := range os.Args[1] {
		if w1 == []rune(os.Args[2])[0] {
			w1 = []rune(os.Args[3])[0]
		}
		z01.PrintRune(w1)
	}
	z01.PrintRune('\n')
}
