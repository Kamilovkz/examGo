package main

import (
	"fmt"
	"strconv"

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
	var result int = n
	for i := 0; i < len(a); i++ {
		result = f(result, a[i])
	}
	for _, ch := range strconv.Itoa(result) {
		z01.PrintRune(ch)
	}
	z01.PrintRune(10)
}
