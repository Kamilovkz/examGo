package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	var first_l bool
	// For first word
	for i := 0; i < len(word); i++ {
		if word[0] == 'a' || word[0] == 'A' || word[0] == 'e' || word[0] == 'E' || word[0] == 'i' || word[0] == 'I' || word[0] == 'o' || word[0] == 'O' || word[0] == 'u' || word[0] == 'U' {
			word += "ay"
			first_l = true
			break
		}
	}
	// for second word
	if !first_l {
		for i := 1; i < len(word); i++ {
			if word[i] == 'a' || word[i] == 'A' || word[i] == 'e' || word[i] == 'E' || word[i] == 'i' || word[i] == 'I' || word[i] == 'o' || word[i] == 'O' || word[i] == 'u' || word[i] == 'U' {
				first_letter := string(word[0])
				word = word[i:]
				word += first_letter
				word += "ay"
				break
			}
		}
	}
	if word == os.Args[1] {
		fmt.Println("No vowels")
	} else {
		fmt.Println(word)
	}
}
