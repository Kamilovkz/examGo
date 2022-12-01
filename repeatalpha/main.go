package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	alpha1 := " abcdefghijklmnopqrstuvwxyz"
	alpha2 := " ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := os.Args[1]
	var count int

	for _, r := range s {
		for i := range alpha1 {
			if r == rune(alpha1[i]) {
				count = i
			}
		}
		for i := range alpha2 {
			if r == rune(alpha2[i]) {
				count = i
			}
		}
		if count != 0 {
			for i := 0; i < count; i++ {
				z01.PrintRune(r)
			}
		} else {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
