package main

import (
	"fmt"
)

func main() {
	input := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(input))
}

func moveZeroes(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for j < len(nums) {
		nums[j] = 0
		j++
	}

	return nums
}
