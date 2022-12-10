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
	if arg[0] == "" {
		z01.PrintRune('\n')
		z01.PrintRune('\n')
		return
	} else {
		slice := Split(arg[0], " ")
		var res string
		for i := len(slice) - 1; i >= 0; i-- {
			res += slice[i] + " "
		}
		for _, word := range res {
			z01.PrintRune(word)
		}
	}
	z01.PrintRune('\n')
}

func Split(s, sep string) []string {
	l := len(sep)
	var result []string
	for i := 0; i < len(s)-l-1; i++ {
		if s[i:i+l] == sep {
			if i < len(s)-l {
				result = append(result, s[:i])
				s = s[i+1:]
				i = 0
			}
		}
	}
	result = append(result, s)
	return result
}
