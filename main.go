package main

import (
	"fmt"
	"hackerrank/warmup"
	"os"
)

func main() {
	args := os.Args[1:]

	problemString := args[0]

	if problemFunc, ok := warmup.Problems[problemString]; ok {
		problemFunc()
	} else {
		fmt.Printf("%s problem not found", problemString)
	}
}
