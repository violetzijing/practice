package main

import "fmt"

func main() {
	input := 11
	fmt.Println(hasAlternatingBits(input))
}

func hasAlternatingBits(n int) bool {
	return n < 3 || n%2 != n/2%2 && hasAlternatingBits(n/2)
}
