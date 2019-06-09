package main

import "fmt"

func main() {
	input := []int{3, 3, 7, 7, 10, 11, 11}
	fmt.Println(singleNonDuplicate(input))
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	for left < right {
		mid = left + (right-left)/2
		if mid%2 == 1 {
			mid--
		}

		if nums[mid] != nums[mid+1] {
			right = mid
		} else {
			left = mid + 2
		}
	}

	return nums[left]
}
