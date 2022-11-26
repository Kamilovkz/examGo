package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	for i, j := 0, 0; i < len(arg[0]) && j < len(arg[1]); {
		if arg[0][i] == arg[1][j] {
			i++
			if i == len(arg[0]) {
				for _, word := range arg[0] {
					z01.PrintRune(word)
				}
				z01.PrintRune('\n')
			}
		}
		j++
		if j == len(arg[1]) {
			return
		}
	}
}
