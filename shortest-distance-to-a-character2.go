package main

func main() {
	input := "loveleetcode"
	found := byte('e')
	shortestToChar(shortestToChar(input, found))
}

func shortestToChar(S string, C byte) []int {
	if len(S) == 0 {
		return []int{}
	}

	left := 0
	right := 0
	result := make([]int, len(S))

	// find right e
	for S[right] != C {
		right++
	}

	// right -> e
	for i := left; i < right; i++ {
		result[i] = right - i
	}
	// end first

}
