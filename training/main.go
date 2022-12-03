package main

import (
	"fmt"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if len(slice) == 0 {
		fmt.Println(slice)
		return
	} else if size == 0 {
		fmt.Println()
		return
	}
	var res [][]int
	for len(slice) > size {
		res = append(res, slice[:size])
		slice = slice[size:]
	}
	res = append(res, slice)
	fmt.Println(res)
}
