package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name string
	Age  int
}

func main() {
	a := &A{}
	v := reflect.ValueOf(a).Elem().FieldByName("Name")
	v2 := reflect.ValueOf(a).Elem().FieldByName("Age")
	baka := reflect.ValueOf("baka")
	age := reflect.ValueOf(11)
	if v.IsValid() {
		v.Set(baka)
	}

	if v2.IsValid() {
		v2.Set(age)
	}

	fmt.Println("a: ", a)
}
