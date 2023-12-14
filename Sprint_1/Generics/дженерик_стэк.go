package main

type Stack[T any] struct {
	T T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.T = val
}

func (s *Stack[T]) Pop() T {
	return s.T
}
