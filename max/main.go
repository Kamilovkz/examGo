package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93, 2222}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
