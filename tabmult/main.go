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
	for i := '1'; i <= '9'; i++ {
		z01.PrintRune(i)
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		for _, i := range os.Args[1] {
			z01.PrintRune(i)
		}
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		printAtoi(os.Args[1], i)
		z01.PrintRune('\n')
	}
}

func printAtoi(s string, r rune) {
	number, _ := strconv.Atoi(s)
	res := number * int(r-48)
	var result []rune
	for res != 0 {
		result = append(result, rune(res%10+48))
		res /= 10
	}
	for d := len(result) - 1; d >= 0; d-- {
		z01.PrintRune(result[d])
	}
}
