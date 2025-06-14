package main

import (
	"fmt"
)

func main() {
	var A, B, C, X int
	fmt.Scanf("%d ", &A)
	fmt.Scanf("%d ", &B)
	fmt.Scanf("%d ", &C)
	fmt.Scanf("%d ", &X)

	count := 0
	for i := 0; i < A+1; i++ {
		for j := 0; j < B+1; j++ {
			for k := 0; k < C+1; k++ {
				sum := i*500 + j*100 + k*50
				if sum == X {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
