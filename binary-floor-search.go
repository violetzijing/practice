package main

import "fmt"

func main() {
	input := []int{-1, 2, 3, 5, 6, 8, 9, 10}
	find := 7
	fmt.Println(BinaryFloorSearch(input, find))
}

func BinaryFloorSearch(arr []int, val int) int {
	var mid int
	left := 0
	right := len(arr) - 1

	for right-left > 1 {
		mid = left + (right-left)/2
		if arr[mid] == val {
			return mid
		}
		if arr[mid] >= val {
			right = mid
		} else {
			left = mid
		}
	}

	return arr[left]
}
