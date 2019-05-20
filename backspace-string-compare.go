package main

import "fmt"

func main() {
	S := "ab##"
	T := "c#d#"
	fmt.Println(backspaceCompare(S, T))
}

func backspaceCompare(S string, T string) bool {
	stack1 := []rune{}
	stack2 := []rune{}

	for _, item := range S {
		if item != '#' {
			stack1 = append(stack1, item)
		} else if len(stack1) != 0 {
			stack1 = stack1[:len(stack1)-1]
		}
	}

	for _, item := range T {
		if item != '#' {
			stack2 = append(stack2, item)
		} else if len(stack2) != 0 {
			stack2 = stack2[:len(stack2)-1]
		}
	}
	fmt.Println("stack1: ", string(stack1))
	fmt.Println("stack2: ", string(stack2))

	if len(stack1) != len(stack2) {
		return false
	}

	return string(stack1) == string(stack2)
}
