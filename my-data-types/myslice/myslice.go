package myslice

// I guess I got stuck here and forgot to commit after googling a bunch of stuff
// was probably about dynamic array allocation looking at the code
// mb wanted to make my custom array to use for stack idk

type MySlice[T any] struct {
	arr   []T
	lastI int
}

func (s *MySlice[T]) Append(item T) {

}

func New[T any]() *MySlice[T] {
	return &MySlice[T]{
		arr:   make([]T, 5),
		lastI: 0,
	}
}
