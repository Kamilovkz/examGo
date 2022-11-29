package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(os.Args) == 1 {
		z01.PrintRune(10)
		return
	}
	fmt.Println(arg)
}
