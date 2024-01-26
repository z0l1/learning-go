package main

import "golang.org/x/exp/constraints"

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Iterater
type Iterater[T any] interface {
	IterReset()
	// IterNext the first IterNext() after IterReset() returns the first value
	IterNext() bool
	IterGet() T
	IterErr() error
}

func IterateData[TStructure Iterater[TItem], TItem any](s TStructure, cb func(TItem)) error {
	s.IterReset()
	for s.IterNext() {
		value := s.IterGet()
		cb(value)
	}

	return s.IterErr()
}
