package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var S, T string
	var pre rune
	is := true

	fmt.Scanf("%s ", &S)
	fmt.Scanf("%s ", &T)

	for i, c := range S {
		if i != 0 && unicode.IsUpper(c) && !strings.ContainsRune(T, pre) {
			is = false
			break
		}
		pre = c
	}

	if is {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
