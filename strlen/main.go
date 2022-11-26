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
	var length int
	for i := 0; i < len(s); i++ {
		length += 1
	}
	return length
}

// OR
// func StrLen(s string) int {
// return len(s)
// }
