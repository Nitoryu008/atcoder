package main

import (
	"fmt"
)

func main() {
	var N, A, B, total int
	fmt.Scanf("%d %d %d", &N, &A, &B)

	for i := 1; i < N+1; i++ {
		sum := i%10 + i%100/10 + i%1000/100 + i%10000/1000 + i/10000
		if A <= sum && sum <= B {
			total += i
		}
	}
	fmt.Println(total)
}
