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
	if len(stones) == 1 {
		return stones[0]
	}
	if len(stones) == 2 && stones[0] != stones[1] {
		return stones[0]
	}
	if len(stones) == 0 {
		return 0
	}
	size := len(stones)
	sort.Ints(stones)
	if stones[size-1] == stones[size-2] {
		stones = stones[:size-2]
		return lastStoneWeight(stones)
	} else if stones[size-1] != stones[size-2] {
		stones[size-2] = stones[size-1] - stones[size-2]
		stones = stones[:size-1]
		return lastStoneWeight(stones)
	}

	return 0
}
