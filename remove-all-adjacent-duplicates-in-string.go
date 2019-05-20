package main

import "fmt"

func main() {
	input := "abbaca"
	fmt.Println(removeDuplicates2(input))
}

func removeDuplicates(S string) string {
	if len(S) == 0 || len(S) == 1 {
		return S
	}
	if len(S) == 2 {
		if S[0] == S[1] {
			return ""
		} else {
			return S
		}
	}
	i := 0
	j := 1
	yaStr := []byte{}
	for j < len(S) {
		if S[i] != S[j] {
			yaStr = append(yaStr, byte(S[i]))
		} else {
			i += 2
			j += 2
			continue
		}
		i++
		j++
	}
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	if i < len(S) {
		yaStr = append(yaStr, byte(S[i]))
	}

	fmt.Println("yaStr: ", string(yaStr))
	if len(S) != len(yaStr) {
		return removeDuplicates(string(yaStr))
	}

	return S
}

func removeDuplicates2(S string) string {
	stack := []byte{}

	for i := 0; i < len(S); i++ {
		if len(stack) != 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}

	return string(stack)
}
