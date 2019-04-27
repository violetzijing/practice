package main

import "fmt"

func main() {
	input := 0
	fmt.Println(fib(input))
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	f0 := 0
	f1 := 1
	if N < 2 {
		return f0 + f1
	}
	var n = 2
	result := f0 + f1
	for n < N {
		f0 = f1
		f1 = result
		result = f0 + f1
		n++
	}

	return result
}
