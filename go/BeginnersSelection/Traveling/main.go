package main

import (
	"fmt"
)

func main() {
	var point_count int
	var points [][]int
	fmt.Scanf("%d ", &point_count)
	points = make([][]int, point_count+1, point_count+1)
	points[0] = []int{0, 0, 0}
	for i := 1; i < point_count+1; i++ {
		points[i] = make([]int, 3, 3)
		fmt.Scanf("%d %d %d ", &points[i][0], &points[i][1], &points[i][2])
	}

	possible := true
	for i := 0; i < point_count; i++ {
		gap := abs(points[i][1]-points[i+1][1]) + abs(points[i][2]-points[i+1][2])
		time := points[i+1][0] - points[i][0]
		if time-gap >= 0 && (time-gap)%2 == 0 {
			continue
		} else {
			possible = false
			break
		}
	}

	if possible {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
