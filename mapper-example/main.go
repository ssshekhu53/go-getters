package main

import (
	"fmt"
	"go-getters/mapper-example/mapper"
)

func main() {
	intArr := make([]int, 0)

	for i := 1; i <= 5; i++ {
		intArr = append(intArr, i)
	}

	fmt.Println("Elements before mapping:")
	for i := range intArr {
		fmt.Printf("intArr[%d]: %d\n", i, intArr[i])
	}

	mapper.Mapper(intArr, func(ele int) int {
		return ele * 2
	})

	fmt.Println("\n\nElements after mapping:")
	for i := range intArr {
		fmt.Printf("intArr[%d]: %d\n", i, intArr[i])
	}
}
