package main

import "fmt"

func main() {
	input := []int{1, 0, 1}
	fmt.Println(singleNumber(input))
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	xor := 0
	for _, n := range nums {
		xor = xor ^ n
	}

	return xor
}
