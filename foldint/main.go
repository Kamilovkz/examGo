package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	Add := func(a int, b int) int {
		return a + b
	}
	Mul := func(a int, b int) int {
		return a * b
	}
	Sub := func(a int, b int) int {
		return a - b
	}
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i < len(a); i++ {
		n = f(n, a[i])
	}
	for _, d := range Itoa(n) {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	neg := ""
	str := ""
	if n < 0 {
		neg = "-"
		n = -n
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		for n > 0 {
			str = string(rune(n%10+48)) + str
			n /= 10
		}
	}
	return neg + str
}
