package stack

type Stack[T any] struct {
	data []T
	topI int
}

// idk but i don't care right now
//func (s *Stack[T]) Push(item T) {
//	s.data[100] =
//	s.data = append(s.data, item)
//
//	s.topI++
//}

func (s *Stack[T]) Pop() (result T, ok bool) {
	if s.topI == -1 {
		return
	}

	result = s.data[s.topI]
	s.topI--
	return
}

func New[T any]() Stack[T] {
	arr := make([]T, 10)
	return Stack[T]{
		data: arr,
		topI: -1,
	}
}
