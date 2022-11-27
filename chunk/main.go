package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	} else if len(slice) == 0 {
		fmt.Println(slice)
		return
	}
	var new_array [][]int
	for len(slice) > size {
		new_array = append(new_array, slice[:size])
		slice = slice[size:]
	}
	new_array = append(new_array, slice)
	fmt.Println(new_array)

}
