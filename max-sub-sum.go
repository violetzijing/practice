package main

import "fmt"

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubSum(input))
}

func maxSubArray(arr []int) int {
	size := len(arr)
	maxEnding := arr[0]
	maxSoFar := arr[0]
	for i := 1; i < size; i++ {
		maxEnding = max(arr[i], arr[i]+maxEnding)
		maxSoFar = max(maxSoFar, maxEnding)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
