package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	hex, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		Printing("ERROR")
		return
	}
	Printing(strconv.FormatInt(hex, 16))
}

func Printing(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
