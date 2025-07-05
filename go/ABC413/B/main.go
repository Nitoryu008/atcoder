package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d ", &N)

	strs := make([]string, N, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s ", &strs[i])
	}

	var result []string
	count := 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			str := strs[i] + strs[j]

			isNew := true
			for _, v := range result {
				if v == str {
					isNew = false
					break
				}
			}
			if isNew {
				result = append(result, str)
				count++
			}
		}
	}

	fmt.Println(count)
}
