package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else if os.Args[1] == "" {
		z01.PrintRune('\n')
		return
	}
	// Not finished
}
