package main

import (
	"fmt"
	"hackerrank/warmup"
	"os"
)

var solutions = []string{
	"repeated_strings",
}

func main() {
	args := os.Args[1:]

	problem := args[0]

	for _, solution := range solutions {
		if problem == solution {
			warmup.RepeatedStrings()
			break
		}
		fmt.Println("problem not found")
	}
}
