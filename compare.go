package main

import (
	"fmt"

	"gopkg.in/volatiletech/null.v6"
)

type A struct {
	Name string
	age  int
}

func main() {
	a1 := A{"baka", 233}
	a2 := A{"baka", 233}
	a3 := A{"233", 233}
	fmt.Println(a1 == a2)
	fmt.Println(a2 == a3)
	fmt.Println("=====")

	num1 := null.NewInt64(1, false)
	num2 := null.Int64From(1)
	num3 := null.Int64{Valid: false}

	fmt.Println(num1 == num2)
	fmt.Println(num1 == num3)
}
