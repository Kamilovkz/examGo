package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}
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
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
