package main

import (
	"fmt"
	"log"

	"go-getters/stack-example/stack"
)

func main() {
	// integer stack
	intStack := stack.New[int]()

	// Pushing data in integer stack
	for i := 1; i <= 5; i++ {
		intStack.Push(i)
	}

	fmt.Printf("Is stack empty: %v\n\n", intStack.IsEmpty())

	// Retrieving data from integer stack
	for i := 1; i <= 5; i++ {
		data, err := intStack.Pop()
		if err != nil {
			log.Fatalf("Some error occurred: %v", err)
		}

		fmt.Printf("Value: %v\n", data)
	}

	fmt.Printf("\nIs stack empty: %v", intStack.IsEmpty())
}
