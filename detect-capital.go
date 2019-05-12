package main

import "fmt"

func main() {
	input := "ffffffffffffffffffffF"
	fmt.Println(detectCapitalUse(input))
}

func detectCapitalUse(word string) bool {
	if len(word) == 0 || len(word) == 1 {
		return true
	}
	startCapital := false
	if word[0] >= 'A' && word[0] <= 'Z' {
		startCapital = true
	}
	lowcase := false
	allCapital := false
	for i := 1; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			lowcase = true
		}
		if word[i] >= 'A' && word[i] <= 'Z' {
			allCapital = true
		}
		if allCapital && (word[i] >= 'a' && word[i] <= 'z') {
			return false
		}
		if startCapital && lowcase && (word[i] >= 'A' && word[i] <= 'Z') {
			return false
		}
		if !startCapital && (word[i] >= 'A' && word[i] <= 'Z') {
			return false
		}
	}

	return true
}
