package main

import (
	"github.com/01-edu/z01"
)

func main() {
	sentence := "zYxWvUtSrQpOnMlKjIhGfEdCbA"
	for _, i := range sentence {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
