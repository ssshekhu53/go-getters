package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Mapper[T constraints.Integer | constraints.Float](arr []T, f func(ele T) T) []T {
	for i := range arr {
		arr[i] = f(arr[i])
	}

	return arr
}

func main() {
	intArr := make([]int, 0)

	for i := 1; i <= 5; i++ {
		intArr = append(intArr, i)
	}

	fmt.Println("Elements before mapping:")
	for i := range intArr {
		fmt.Printf("intArr[%d]: %d\n", i, intArr[i])
	}

	Mapper(intArr, func(ele int) int {
		return ele * 2
	})

	fmt.Println("\n\nElements after mapping:")
	for i := range intArr {
		fmt.Printf("intArr[%d]: %d\n", i, intArr[i])
	}
}
