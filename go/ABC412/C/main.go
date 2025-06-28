package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var T int
	var N int
	var result []int
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)

	sc.Scan()
	T, _ = strconv.Atoi(sc.Text())
	result = make([]int, T, T)

	for t := 0; t < T; t++ {
		sc.Scan()
		N, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		var S = make([]int, N, N)
		var used = make([]bool, N, N)
		for i := 0; i < N; i++ {
			S[i], _ = strconv.Atoi(inputs[i])
		}

		count := 1
		last := 0
		for true {
			if 2*S[last] >= S[N-1] {
				count++
				break
			}

			max_i := -1
			for i := 1; i < N; i++ {
				if used[i] {
					continue
				}
				if S[i] <= 2*S[last] {
					if max_i == -1 || S[i] > S[max_i] {
						max_i = i
					}
				}
			}

			if max_i == -1 || S[max_i] <= S[last] {
				count = -1
				break
			}

			count++
			used[max_i] = true
			last = max_i
		}

		result[t] = count
	}

	for i := 0; i < T; i++ {
		fmt.Println(result[i])
	}
}
