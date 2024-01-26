package main

type LinkedListNode[T any] struct {
	Next *LinkedListNode[T]
	Data T
}

type LinkedList[T any] struct {
	head   *LinkedListNode[T]
	length uint
}

func (ll *LinkedList[T]) Len() uint {
	return ll.length
}

func (ll *LinkedList[T]) Add(item T) {
	newNode := &LinkedListNode[T]{
		Next: nil,
		Data: item,
	}
	ll.length++

	if ll.head == nil {
		ll.head = newNode
		return
	}

	iterNode := ll.head
	for iterNode.Next != nil {
		iterNode = iterNode.Next
	}

	iterNode.Next = newNode
}

func (ll *LinkedList[T]) Get(idx uint) (T, bool) {
	var nilItem T
	if idx >= ll.length {
		return nilItem, false
	}

	iterNode := ll.head
	for nodeI := uint(0); nodeI < idx; nodeI++ {
		iterNode = iterNode.Next
	}

	return iterNode.Data, true
}

func (ll *LinkedList[T]) Remove(idx uint) bool {
	if idx >= ll.length {
		return false
	}

	if idx == 0 {
		ll.head = ll.head.Next
		return true
	}

	prev := ll.head
	iterNode := ll.head.Next
	for nodeI := uint(1); nodeI < idx; nodeI++ {
		prev = iterNode
		iterNode = iterNode.Next
	}

	prev.Next = iterNode.Next
	return true
}

func (ll *LinkedList[T]) ChanIter(buffered bool) <-chan T {
	var ch chan T

	if buffered {
		ch = make(chan T, ll.length)
	} else {
		ch = make(chan T)
	}

	go func() {
		iterNode := ll.head
		for iterNode != nil {
			ch <- iterNode.Data
			iterNode = iterNode.Next
		}

		close(ch)
	}()

	return ch
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{
		head:   nil,
		length: 0,
	}
}
