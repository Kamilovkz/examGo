package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := []byte(os.Args[1])
	size := len(word)
	if !checkConstant(word) {
		Printer("No vowels")
		return
	}
	result := string(addAy(word))
	result = string(addAy(pushConsonant(word, size)))
	Printer(result)
}

func checkConstant(word []byte) bool {
	var vowel bool
	for _, b := range word {
		if checkVowel(b) {
			vowel = true
		}
	}
	return vowel
}

func Printer(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}

func pushConsonant(word []byte, size int) []byte {
	for i, b := range word {
		if !checkVowel(b) {
			word = append(word, b)
			word = word[i:]
		} else {
			break
		}
	}
	return word[(len(word) - size):]
}

func addAy(word []byte) []byte {
	if checkVowel(word[0]) {
		word = append(word, 'a', 'y')
	}
	return word
}

func checkVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
