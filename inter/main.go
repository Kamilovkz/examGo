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
	for _, word := range result {
		z01.PrintRune(word)
	}
	z01.PrintRune(10)

}

func double(s string, r rune) bool {
	for _, i := range s {
		if i == r {
			return true
		}
	}
	return false
}
