package main

import "fmt"

func main() {
	input := []int{2, 5, 4, 6, 1, 3}
	fmt.Println(selectionSort(input))
}

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		minIndex := i
		for j := i; j < len(nums); j++ {
			if min > nums[j] {
				min = nums[j]
				minIndex = j
			}
		}
		tmp := nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = tmp
		fmt.Println("nums: ", nums)
	}
	return nums
}
