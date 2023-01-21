package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	var result string
	for _, i := range os.Args[1] {
		for _, j := range os.Args[2] {
			if i == j && !double(result, i) {
				result += string(i)
			}
		}
	}
	for _, d := range result {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}

func double(s string, r rune) bool {
	for _, d := range s {
		if d == r {
			return true
		}
	}
	return false
}
