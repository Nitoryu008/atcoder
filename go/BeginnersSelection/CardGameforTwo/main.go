package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var cards []int
	var card_count, alice, bob int

	fmt.Scanf("%d ", &card_count)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	for i := 0; i < card_count; i++ {
		num, _ := strconv.Atoi(inputs[i])
		cards = append(cards, num)
	}
	sort.Slice(cards, func(i, j int) bool { return cards[i] < cards[j] })
	for i := 0; i < card_count; i++ {
		if i%2 == 0 {
			alice += cards[card_count-i-1]
		} else {
			bob += cards[card_count-i-1]
		}
	}
	fmt.Println(alice - bob)
}
