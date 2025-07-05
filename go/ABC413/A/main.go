package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, M int
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	N, _ = strconv.Atoi(inputs[0])
	M, _ = strconv.Atoi(inputs[1])

	sc.Scan()
	inputs = strings.Split(sc.Text(), " ")

	total := 0
	for i := 0; i < N; i++ {
		weight, _ := strconv.Atoi(inputs[i])
		total += weight
	}

	if total <= M {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}
