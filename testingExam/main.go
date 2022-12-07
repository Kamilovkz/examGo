package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d: ", i)
		randomizer()
	}
}

func randomizer() {
	rand.Seed(time.Now().UnixNano())
	charset := "123456789"
	c := charset[rand.Intn(len(charset))]
	fmt.Println(string(c))
}
