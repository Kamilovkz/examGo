package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res := ""
		word := os.Args[1]
		for i := len(word) - 1; i >= 0; i-- {
			if word[i] != ' ' {
				res = string(word[i]) + res
			} else if word[i] == ' ' && res != "" {
				break
			}
		}
		if res == "" {
			return
		}
		for _, d := range res {
			z01.PrintRune(d)
		}
	}
	z01.PrintRune('\n')
}
