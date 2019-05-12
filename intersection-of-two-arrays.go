package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	result := []int{}
	numsMap := map[int]int{}
	for _, item := range nums1 {
		numsMap[item] = 1
	}
	for _, item := range nums2 {
		if numsMap[item] == 1 {
			numsMap[item]++
			result = append(result, item)
		}
	}

	return result
}
