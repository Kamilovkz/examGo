package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	if checkPrime(nb) {
		return nb
	}
	if !checkPrime(nb) {
		for {
			nb--
			if checkPrime(nb) {
				return nb
			}
		}
	}
	return nb
}

func checkPrime(nb int) bool {
	if nb <= 1 && nb != 0 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
