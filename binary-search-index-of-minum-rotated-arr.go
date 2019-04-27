package main

import "fmt"

func main() {
	input := []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	fmt.Println(BinarySearchIndexOfMinimumRotatedArray(input, 0, 9))
}

func BinarySearchIndexOfMinimumRotatedArray(arr []int, left int, right int) int {
	if arr[left] == arr[right] {
		return left
	}

	var mid int
	for left <= right {
		if left == right {
			return left
		}

		mid = left + (right-left)/2
		if arr[mid] < arr[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}
