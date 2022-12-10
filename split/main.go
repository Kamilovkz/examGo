package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	l := len(sep)
	var result []string
	for i := 0; i < len(s)-l-1; i++ {
		if s[i:i+l] == sep && i < len(s)-1 {
			result = append(result, s[:i])
			s = s[i+l:]
			i = 0
		}
	}
	result = append(result, s)
	var result2 []string
	for _, word := range result {
		if word != "" && word != " " {
			result2 = append(result2, word)
		} else {
			continue
		}
	}
	return result2
}
