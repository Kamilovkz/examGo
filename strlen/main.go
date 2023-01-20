package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

// Our program here

func StrLen(s string) int {
	return len([]rune(s))
}
