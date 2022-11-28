package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num := atoi(os.Args[1])
	if num >= 4000 || num == 0 {
		fmt.Printf("%s\n", "ERROR: cannot convert to roman digit")
		return
	}
}

func atoi(s string) int {
	var negative bool
	var n int
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0
		}
		n = n*10 + int(ch)
	}
	if negative {
		n = -n
	}
	return n
}
