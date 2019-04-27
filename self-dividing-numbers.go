package main

import (
	"fmt"
	"math"
)

func main() {
	left := 1
	right := 22

	fmt.Println(selfDividingNumbers(left, right))
}

func selfDividingNumbers(left int, right int) []int {
	result := []int{}

	for num := left; num <= right; num++ {
		var numSlice = []int{}
		var divided = []int{10000, 1000, 100, 10, 1}
		var flag = true
		dividedNum := num
		for _, d := range divided {
			if num < d {
				continue
			}
			res := dividedNum / d
			if res != 0 {
				numSlice = append(numSlice, res)
				dividedNum = dividedNum - res*d
			} else {
				flag = false
			}
		}
		if !flag {
			continue
		}
		for _, n := range numSlice {
			if math.Mod(float64(num), float64(n)) != 0 {
				flag = false
			}

		}
		if flag {
			result = append(result, num)
		}
	}

	return result
}
