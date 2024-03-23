package stack

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type stack[T constraints.Ordered] struct {
	data []T
}

func New[T constraints.Ordered]() *stack[T] {
	data := make([]T, 0)

	return &stack[T]{data: data}
}

func (s *stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

func (s *stack[T]) Pop() (T, error) {
	n := len(s.data)
	if n == 0 {
		return T(0), errors.New("stack is empty")
	}

	data := s.data[n-1]
	s.data = s.data[0 : n-1]

	return data, nil
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
