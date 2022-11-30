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
	a, _ := strconv.Atoi(os.Args[1])
	var str string
	if a == 0 {
		str = "00000000"
	}
	for len(str) != 8 {
		str = string(rune(a%2+48)) + str
		a /= 2
	}
	str2 := str[4:8] + str[0:4]
	for _, i := range str {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
	for _, i := range str2 {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
