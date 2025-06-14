package main

import (
	"fmt"
)

func main() {
	var s string
	divide := [4]string{"dream", "dreamer", "erase", "eraser"}
	fmt.Scanln(&s)
	s = Reverse(s)
	for i := 0; i < len(divide); i++ {
		divide[i] = Reverse(divide[i])
	}

	possible := true
LOOP:
	for i := 0; i < len(s); {
		for j := 0; j < 4; j++ {
			d := divide[j]
			if len(s) >= i+len(d) && s[i:i+len(d)] == d {
				i += len(d)
				continue LOOP
			}
		}
		possible = false
		break
	}

	if possible {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
