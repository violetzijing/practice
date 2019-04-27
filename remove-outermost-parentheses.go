package main

import "fmt"

func main() {
	input := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(input))
}

func removeOuterParentheses(S string) string {
	leftCount := 0
	output := []rune{}

	leftPar := '('
	rightPar := ')'
	for _, s := range S {
		if s == leftPar {
			if leftCount != 0 {
				output = append(output, leftPar)
			}
			leftCount++
		} else {
			leftCount--
			if leftCount != 0 {
				output = append(output, rightPar)
			}
		}
	}

	return string(output)
}
