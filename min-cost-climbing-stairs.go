package main

import (
	"fmt"
)

func main() {
	input := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(input))
}

func minCostClimbingStairs(cost []int) int {
	size := len(cost)
	dp := make([]int, size)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < size; i++ {
		if dp[i-2] > dp[i-1] {
			dp[i] = cost[i] + dp[i-1]
		} else {
			dp[i] = cost[i] + dp[i-2]
		}
	}

	if dp[size-1] > dp[size-2] {
		return dp[size-2]
	}

	return dp[size-1]
}
