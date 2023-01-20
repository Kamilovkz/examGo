package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	} else {
		res := ""
		for i := len(args[0]) - 1; i >= 0; i-- {
			if args[0][i] != ' ' {
				res = string(args[0][i]) + res
			} else if args[0][i] == ' ' && res != "" {
				break
			}
		}
		if res == "" {
			return
		}
		for _, i := range res {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
