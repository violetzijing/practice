package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{2, 1, 2, 1, 1, 2, 2, 1}
	fmt.Println(heightChecker(input))
}

func heightChecker(heights []int) int {
	originHeights := make([]int, len(heights))
	for index, item := range heights {
		originHeights[index] = item
	}
	sort.Ints(heights)
	count := 0
	for i := 0; i < len(heights); i++ {
		if originHeights[i] != heights[i] {
			count++
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
