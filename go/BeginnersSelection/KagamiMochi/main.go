package main

import (
	"fmt"
	"sort"
)

func main() {
	var mochi_count int
	fmt.Scanf("%d ", &mochi_count)
	mochi_array := make([]int, mochi_count, mochi_count)
	for i := 0; i < mochi_count; i++ {
		fmt.Scanf("%d ", &mochi_array[i])
	}

	sort.Slice(mochi_array, func(i, j int) bool { return mochi_array[i] < mochi_array[j] })
	dup_count := 0
	for i := 0; i < mochi_count-1; i++ {
		if mochi_array[i] == mochi_array[i+1] {
			dup_count += 1
		}
	}
	fmt.Println(mochi_count - dup_count)
}
