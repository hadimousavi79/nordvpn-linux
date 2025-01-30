package internal

import "fmt"

type GenericError[T any] struct {
	Value T
}

func NewGenericError[T any](value T) *GenericError[T] {
	return &GenericError[T]{Value: value}
}

func (e *GenericError[T]) Error() string {
	return fmt.Sprintf("error: %v", e.Value)
}
