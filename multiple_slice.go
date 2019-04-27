package main

import "fmt"

type structA struct {
	ID         int
	ConnectBID int
	Name       string
	DID        int
}

type structB struct {
	ID         int
	ConnectCID int
	Name       string
}

type structC struct {
	ID         int
	ConnectDID int
	Name       string
}

type structD struct {
	ID   int
	Name string
}

func main() {
	as := []*structA{
		&structA{1, 1, "B 1", 0},
		&structA{2, 2, "B 2", 0},
	}

	bs := []*structB{
		&structB{1, 1, "A1, C1"},
		&structB{2, 2, "A2, C2"},
	}

	cs := []*structC{
		&structC{1, 1, "B1 D1"},
		&structC{2, 2, "B2 D2"},
	}

	ds := []*structD{
		&structD{1, "C1"},
		&structD{2, "C2"},
	}
	var a2BMap = map[int]*structB{}
	for _, b := range bs {
		a2BMap[b.ID] = b
	}
	var b2CMap = map[int]*structC{}
	for _, c := range cs {
		b2CMap[c.ID] = c
	}
	var c2DMap = map[int]*structD{}
	for _, d := range ds {
		c2DMap[d.ID] = d
	}

	for _, a := range as {
		b := a2BMap[a.ConnectBID]
		c := b2CMap[b.ConnectCID]
		d := c2DMap[c.ConnectDID]
		anotherD := c2DMap[b2CMap[a2BMap[a.ConnectBID].ConnectCID].ConnectDID]
		fmt.Println("d: ", d)
		fmt.Println("anotherD: ", anotherD)
		a.DID = d.ID
		fmt.Println("a: ", a)
	}
}
