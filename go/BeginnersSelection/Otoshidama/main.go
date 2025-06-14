package main

import (
	"fmt"
)

func main() {
	var cash_count, total_amount int
	fmt.Scanf("%d %d ", &cash_count, &total_amount)

	possible := false
LOOP:
	for i := 0; i < cash_count+1; i++ {
		for j := 0; j < cash_count+1-i; j++ {
			k := cash_count - i - j
			if i*10000+j*5000+k*1000 == total_amount {
				fmt.Printf("%d %d %d\n", i, j, k)
				possible = true
				break LOOP
			}
		}
	}

	if !possible {
		fmt.Printf("-1 -1 -1")
	}
}
