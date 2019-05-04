package main

import "fmt"

func main() {
	input := 14
	fmt.Println(fizzBuzz(input))
}

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		result[i-1] = fmt.Sprintf("%d", i)
	}

	for i := 2; i < n; i += 3 {
		fmt.Println("i: ", i)
		result[i] = "Fizz"
	}

	for i := 4; i < n; i += 5 {
		result[i] = "Buzz"
	}

	for i := 14; i < n; i += 15 {
		result[i] = "FizzBuzz"
	}

	return result
}
