package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	find := 55
	fmt.Println(JumpSearch(input, find))
}

func JumpSearch(s []int, val int) int {
	n := len(s)
	step := math.Sqrt(float64(n))
	prev := 0
	for s[int(math.Min(step, float64(n)))] < val {
		prev = int(step)
		step += math.Sqrt(float64(n))
		if prev > n {
			return -1
		}
	}

	for s[prev] < val {
		prev++
		if prev == int(math.Min(step, float64(n))) {
			return -1
		}
	}

	if s[prev] == val {
		return prev
	}

	return -1
}
