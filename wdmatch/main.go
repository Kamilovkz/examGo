package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argm := os.Args[1:]
	if len(argm) != 2 {
		return
	}
	for i, j := 0, 0; i < len(argm[0]) && j < len(argm[1]); {
		if argm[0][i] == argm[1][j] {
			i++
			if i == len(argm[0]) {
				for _, word := range argm[0] {
					z01.PrintRune(word)
				}
				z01.PrintRune('\n')
			}
		}
		j++
		if j == len(argm[1]) {
			return
		}
	}
}
