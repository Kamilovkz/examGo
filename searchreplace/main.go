package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	words := os.Args[1:]
	for _, word1 := range words[0] {
		for _, word2 := range words[1] {
			for _, word3 := range words[2] {
				if word1 == word2 {
					word1 = word3
				}
			}
		}
	}
	for _, d := range words[0] {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}

// go run . "hallo thara" "a" "e"
