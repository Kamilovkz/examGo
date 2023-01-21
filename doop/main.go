package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		arg := os.Args[1:]
		a := Atoi(arg[0])
		b := Atoi(arg[2])
		if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
			return
		}
		if len(arg[0]) >= 19 {
			return
		}
		switch arg[1] {
		case "+":
			fmt.Printf("%d\n", a+b)
		case "-":
			fmt.Printf("%d\n", a-b)
		case "*":
			fmt.Printf("%d\n", a*b)
		case "/":
			if b == 0 {
				fmt.Printf("No division by 0\n")
				return
			}
			fmt.Printf("%d\n", a/b)
		case "%":
			if b == 0 {
				fmt.Printf("No modulo by 0\n")
				return
			}
			fmt.Printf("%d\n", a%b)
		}
	}
}

func Atoi(str string) int {
	neg := 1
	if str[0] == '-' {
		neg = 01
		str = str[1:]
	}
	n := 0
	for _, w := range str {
		if w < '0' || w > '9' {
			os.Exit(0)
		}
		n = n*10 + int(w-'0')
	}
	return neg * n
}
