package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	number := os.Args[1:]
	if len(number) != 1 {
		z01.PrintRune('\n')
		return
	}
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
		z01.PrintRune('\n')
	}
}
