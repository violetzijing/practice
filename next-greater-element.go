package main

import "fmt"

func main() {
	input := []int{5, 3, 4, 9}
	fmt.Println(NextGreatElement(input))
}

func NextGreatElement(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return []int{-1}
	}

	result := make([]int, len(nums))
	if nums[0] > nums[len(nums)-1] {
		result[len(nums)-1] = nums[0]
	} else {
		result[len(nums)-1] = -1
	}
	currentMax := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if currentMax > nums[i+1] && nums[i+1] > nums[i] {
			currentMax = nums[i+1]
		}
		if currentMax > nums[i] {
			result[i] = currentMax
		} else {
			result[i] = -1
		}
		fmt.Println("nums[i+1]: ", nums[i+1])
		fmt.Println("nums[i]: ", nums[i])
		fmt.Println("currentMax: ", currentMax)
		fmt.Println("result[i]: ", result[i])
		fmt.Println("====")
	}

	return result
}
