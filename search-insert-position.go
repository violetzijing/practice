package main

import "fmt"

func main() {
	input := []int{1, 3, 4, 6}
	found := 40
	fmt.Println(searchInsert(input, found))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	for left < right {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		mid = left + (right-left)/2
	}

	if nums[mid] < target {
		return mid + 1
	}

	return mid
}
