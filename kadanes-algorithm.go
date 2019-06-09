package main

import "fmt"

func main() {
	input := []int{1, 2, 3, -2, 5}
	fmt.Println(kadaneAlgo(input))
}

func kadaneAlgo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], dp[i-1])
	}

	max := dp[0]
	for _, item := range dp {
		if item > max {
			max = item
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
