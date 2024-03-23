package mapper

import "golang.org/x/exp/constraints"

func Mapper[T constraints.Integer | constraints.Float](arr []T, f func(ele T) T) []T {
	for i := range arr {
		arr[i] = f(arr[i])
	}

	return arr
}
