package main

import "fmt"

func main() {
	input := "III"
	fmt.Println(diStringMatch(input))
}

func diStringMatch(S string) []int {
	result := make([]int, len(S)+1)
	increase := 0
	decrease := len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			result[i] = increase
			increase++
		}
		if S[i] == 'D' {
			result[i] = decrease
			decrease--
		}
	}
	result[len(S)] = decrease

	return result
}
