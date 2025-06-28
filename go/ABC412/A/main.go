package main

import (
	"fmt"
)

func main() {
	var N int
	var A, B int
	var count int = 0

	fmt.Scanf("%d ", &N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d ", &A, &B)
		if A < B {
			count++
		}
	}

	fmt.Print(count)
}
