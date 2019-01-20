// https://leetcode.com/problems/unique-morse-code-words/submissions/

package main

import (
	"fmt"
)

func main() {

	test := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(test))

	fmt.Println("Hello, playground")
}

func uniqueMorseRepresentations(words []string) int {
	var morseMap = map[rune]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}

	var result = map[string]int{}
	for _, word := range words {
		str := ""
		for _, w := range word {
			str += morseMap[w]
		}
		if _, ok := result[str]; !ok {
			result[str] = 1
		} else {
			result[str] += 1
		}
	}

	return len(result)
}

