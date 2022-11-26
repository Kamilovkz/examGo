package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	f := os.Args[1]
	for _, i := range f {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
