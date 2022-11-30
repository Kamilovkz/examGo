package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	if n1 < n2 {
		n1, n2 = n2, n1
	}
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	fmt.Println(n1)
}

// My solution
// if len(os.Args) != 3 {
// 	return
// }
// else if len(os.Args) > 3 {
// 	fmt.Println()
// 	return
// }
// arg := os.Args[1:]
// firstNum, _ := strconv.Atoi(arg[0])
// secondNumb, _ := strconv.Atoi(arg[1])
// max_number := 0
// if firstNum > secondNumb {
// 	max_number = firstNum
// } else {
// 	max_number = secondNumb
// }
// for i := max_number; i >= 0; i-- {
// 	if firstNum%i == 0 && secondNumb%i == 0 {
// 		fmt.Println(i)
// 		break
// 	}
// }
