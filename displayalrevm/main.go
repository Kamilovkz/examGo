package main

import (
	"github.com/01-edu/z01"
)

func main() {
	s := 'Z'
	for i := 'z'; i > 'a'; i-- {
		if i%2 == 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(s)
		}
		s--
	}
	z01.PrintRune('\n')
}
