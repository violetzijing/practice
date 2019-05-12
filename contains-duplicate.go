package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(input))
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	numMap := map[int]bool{}
	for _, item := range nums {
		if numMap[item] {
			return true
		}
		numMap[item] = true
	}

	return false
}
