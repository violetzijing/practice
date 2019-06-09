package main

import "fmt"

func main() {
	input := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(input))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	d := make([]int, n)
	d[0] = 1
	for i := 1; i < n; i++ {
		longest := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				candidate := d[j] + 1
				if candidate > longest {
					longest = candidate
				}
			}
		}
		d[i] = longest
	}

	max := d[0]
	for _, item := range d {
		if max < item {
			max = item
		}
	}

	return max
}
