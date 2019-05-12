package main

import "fmt"

func main() {
	input := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(insertionSort(input))
}

func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
		fmt.Println("nums: ", nums)
	}

	return nums
}
