package main

import "github.com/guregu/null"

func main() {
	type A struct {
		Name *null.String
	}
	// wrong
	a := &A{
		Name: &null.String("baka"),
	}

	// correct
	name := null.String("baka")
	a := &A{
		Name: &name,
	}
}
