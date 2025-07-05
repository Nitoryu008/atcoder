package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers [][]int
	var results []int
	current_i := 0
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	query_count, _ := strconv.Atoi(sc.Text())
	head_index := 0

	for i := 0; i < query_count; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		typ, _ := strconv.Atoi(inputs[0])

		count, _ := strconv.Atoi(inputs[1])
		if typ == 1 {
			num, _ := strconv.Atoi(inputs[2])

			numbers = append(numbers, []int{count, num})
		} else {
			total := 0

			for count > 0 {
				c := numbers[head_index][0]
				x := numbers[head_index][1]
				deleted_count := c - current_i
				if deleted_count > count {
					deleted_count = count
					current_i += count
				} else {
					current_i = 0
					head_index++
				}

				total += deleted_count * x
				count -= deleted_count
			}

			results = append(results, total)

			if head_index > 1000 {
				numbers = numbers[head_index:]
				head_index = 0
			}
		}
	}

	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}
}
