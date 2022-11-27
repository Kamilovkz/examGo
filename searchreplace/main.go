package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	}
	for _, word := range arg[0] {
		for _, word1 := range arg[1] {
			for _, word2 := range arg[2] {
				if word == word1 {
					word = word2
				}
			}
		}
		z01.PrintRune(word)
	}

	z01.PrintRune('\n')
}
