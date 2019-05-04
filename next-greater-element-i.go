package main

import "fmt"

func main() {
	num1 := []int{2, 4}
	num2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(num1, num2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	num2Map := map[int]int{}
	result := make([]int, len(nums1))

	for i := 0; i < len(nums2); i++ {
		if i == len(nums2)-1 {
			num2Map[nums2[i]] = -1
			break
		}
		num2Map[nums2[i]] = nums2[i+1]
	}

	for i := 0; i < len(nums1); i++ {
		if nums1[i] < num2Map[nums1[i]] {
			result[i] = num2Map[nums1[i]]
		} else {
			result[i] = -1
		}
	}

	return result
}
