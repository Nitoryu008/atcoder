package main

import "fmt"

func main() {
	var s string
	var count int
	fmt.Scan(&s)
	for _, c := range s {
		if (string(c) == "1") {
			count += 1
		}
	}
	fmt.Print(count)
}