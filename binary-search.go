package main

import "fmt"

func main() {
	input := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	find := 55
	fmt.Println(BinarySearch(input, find))
}

func BinarySearch(arr []int, val int) int {
	var mid int
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] == val {
			return mid
		}
		if arr[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
