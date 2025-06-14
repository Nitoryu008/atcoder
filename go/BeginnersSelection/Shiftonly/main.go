package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	var A []int

	fmt.Scanf("%d ", &N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	for _, input := range inputs {
		p, _ := strconv.Atoi(input)
		A = append(A, p)
	}

	i := 0
LOOP:
	for {
		for j := 0; j < N; j++ {
			if A[j]%2 != 0 {
				break LOOP
			} else {
				A[j] /= 2
			}
		}
		i++
	}

	fmt.Println(i)
}
