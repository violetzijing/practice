package main

import (
	"fmt"

	"gopkg.in/volatiletech/null.v6"
)

type MyStruct struct {
	name *null.String
}

func main() {
	name := null.StringFrom("baka")
	myStruct1 := &MyStruct{
		name: &name,
	}
	fmt.Println("myStruct: ", myStruct1)

	/*
		myStruct2 := &MyStruct{
			name: &null.StringFrom("baka"),
		}
		fmt.Println("myStruct: ", myStruct2)
	*/
}
