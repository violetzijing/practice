package main

import "fmt"

func main() {
	input := "a1b2"
	fmt.Println(letterCasePermutation(input))
}

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{}
	}

}
