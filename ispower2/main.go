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
	n := Atoi(arg[0])
	if n&(n-1) == 0 && n > 0 {
		result := "true"
		for _, w := range result {
			z01.PrintRune(w)
		}
		z01.PrintRune(10)
	} else {
		result := "false"
		for _, w := range result {
			z01.PrintRune(w)
		}
		z01.PrintRune(10)
	}
}

func Atoi(args string) int {
	s := 0
	temp := 1
	if args[0] == '-' {
		temp = -1
		args = args[1:]
	}
	for v := range args {
		s = s * 10
		s = s + int(args[v]-48)
	}
	return s * temp
}
