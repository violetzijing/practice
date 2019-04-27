package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now().UTC()
	time.Sleep(10 * time.Second)
	t2 := time.Now().UTC()

	fmt.Println(t2.Before(t1))
	fmt.Println(t2.After(t1))
}
