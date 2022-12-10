package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1:]
	res := ""
	for i := len(arg[0]) - 1; i >= 0; i-- {
		if arg[0][i] != ' ' {
			res = string(arg[0][i]) + res
		}
		if res != "" && arg[0][i] == ' ' {
			break
		}
	}
	if res == "" {
		return
	}
	for _, word := range res {
		z01.PrintRune(word)
	}
	z01.PrintRune(10)
}
