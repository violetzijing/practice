package main

import "fmt"

func main() {
	input := []int{1, 4, 20, 3, 10, 5}
	fmt.Println(findSubarrWithGivenSum(input, 33))
}

func findSubarrWithGivenSum(arr []int, sum int) []int {
	n := len(arr)
	currSum := arr[0]
	start := 0

	for i := 1; i < n; i++ {
		for currSum > sum && start < i-1 {
			currSum = currSum - arr[start]
			start++
		}

		if currSum == sum {
			return arr[start:i]
		}
		if i < n {
			currSum = currSum + arr[i]
		}
	}

	return []int{}
}
