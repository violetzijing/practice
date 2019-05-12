package main

import "fmt"

func main() {
	input := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(input))
}

func commonChars(A []string) []string {
	if len(A) == 0 {
		return A
	}
	result := []string{}

	strMap := map[rune]map[int]int{}
	for index, str := range A {
		for _, s := range str {
			if _, ok := strMap[s][index]; ok {
				strMap[s][index]++
			} else {
				strMap[s] = map[int]int{
					index: 1,
				}
			}
		}
	}

	for key, value := range strMap {
		count := 0
		if value
	}
}
