package main

import "fmt"

func main() {
	input := []int{2}
	amount := 3
	fmt.Println(coinChange(input, amount))
}

func coinChange(coins []int, amount int) int {
	d := make([]int, amount+1)
	d[0] = 0

	for i := 1; i <= amount; i++ {
		plans := []int{}
		for _, coin := range coins {
			if coin > i {
				continue
			}
			if d[i-coin] != -1 {
				plans = append(plans, d[i-coin]+1)
			}
		}
		if len(plans) == 0 {
			d[i] = -1
			continue
		}
		min := plans[0]
		for _, i := range plans {
			if min > i {
				min = i
			}
		}
		d[i] = min
	}

	return d[amount]
}
