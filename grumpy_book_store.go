package main

import "fmt"

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	fmt.Println(maxSatisfied(customers, grumpy, 3))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	if len(customers) == 0 {
		return 0
	}
	satisfy := make([]int, len(customers))
	for i := 0; i <= len(customers)-X; i++ {
		for j := i; j < i+X; j++ {
			if grumpy[j] == 1 {
				satisfy[i] += customers[j]
			}
		}
	}
	fmt.Println("satisfy: ", satisfy)

	max := satisfy[0]
	maxIndex := 0
	for index, item := range satisfy {
		if max < item {
			max = item
			maxIndex = index
		}
	}
	fmt.Println("maxIndex: ", maxIndex)
	for i := maxIndex; i < maxIndex+X; i++ {
		grumpy[i] = 0
	}
	fmt.Println("grumpy: ", grumpy)
	count := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			count += customers[i]
		}
	}

	return count
}
