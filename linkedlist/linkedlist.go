package main

type LinkedListNode[T any] struct {
	Next *LinkedListNode[T]
	Data T
}

type LinkedList[T any] struct {
	head   *LinkedListNode[T]
	length uint
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
	if idx < 0 || idx > ll.length-1 {
		return nilItem, false
	}

	iterNode := ll.head
	for nodeI := uint(0); nodeI < idx; nodeI++ {
		iterNode = iterNode.Next
	}

	return iterNode.Data, true
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{
		head:   nil,
		length: 0,
	}
}
