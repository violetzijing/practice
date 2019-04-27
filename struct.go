package main

import (
	"encoding/json"
	"fmt"
)

type ConflictDetection struct {
	Result SubResult `json:"result"`
}

type SubResult struct {
	Result string `json:"result"`
	Reason string `json:"reason"`
}

func main() {
	cd := ConflictDetection{
		Result: SubResult{"conflict", "233"},
	}

	b, err := json.Marshal(cd)
	fmt.Println(string(b))
	fmt.Println(err)
}
