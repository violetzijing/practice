package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{1, 1, 4, 2, 2}
	fmt.Println(lastStoneWeight(input))
}

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)

	for len(stones) > 1 {
		x := stones[len(stones)-1]
		y := stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if x != y {
			stones = append(stones, x-y)
		}
		fmt.Println("stones: ", stones)
		sort.Ints(stones)
		if len(stones) == 1 {
			break
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}
