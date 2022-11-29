package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num := atoi(os.Args[1])
	if num >= 4000 || num == 0 {
		fmt.Printf("%s\n", "ERROR: cannot convert to roman digit")
		return
	}
	ones1 := []string{"", "I+", "I+I+", "I+I+I+", "(V-I)+", "V+", "V+I+", "V+I+I+", "V+I+I+I+", "(X-I)+"}
	tens1 := []string{"", "X+", "X+X+", "X+X+X+", "(L-X)+", "L+", "L+X+", "L+X+X+", "L+X+X+X+", "(C-X)+"}
	hundreds1 := []string{"", "C+", "C+C+", "C+C+C+", "(D-C)+", "D+", "D+C+", "D+C+C+", "D+C+C+C+", "(M-C)+"}
	thousands1 := []string{"", "M+", "M+M+", "M+M+M+"}
	res1 := thousands1[num/1000] + hundreds1[num%1000/100] + tens1[num%100/10] + ones1[num%10]
	res1 = res1[:len(res1)-1]
	fmt.Printf("%s\n", res1)

	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands := []string{"", "M", "MM", "MMM"}
	res := thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
	fmt.Printf("%s\n", res)
}

func atoi(s string) int {
	var negative bool
	var n int
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0
		}
		n = n*10 + int(ch)
	}
	if negative {
		n = -n
	}
	return n
}
