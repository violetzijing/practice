package main

import "fmt"

type MyStruct struct {
	*Name
	MyName string
}

type Name struct {
	MyName string
}

func test() {
	a := &MyStruct{
		MyName: "baka1",
	}

	b := &MyStruct{
		Name: &Name{MyName: "baka2"},
	}

	c := &MyStruct{
		Name:   &Name{MyName: "baka3"},
		MyName: "baka4",
	}

	fmt.Println("a.MyName: ", a.MyNam)
	fmt.Println("b.MyName: ", b.MyName)
	fmt.Println("c.MyName: ", c.MyName)
}
