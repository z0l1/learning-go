package main

import "fmt"

func main() {

	ll := NewLinkedList[int]()

	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	//fmt.Println(ll.Get(-1))
	fmt.Println(ll.Get(0))
	fmt.Println(ll.Get(1))
	fmt.Println(ll.Get(2))
	fmt.Println(ll.Get(3))

	//for i := range ll {
	//
	//}

}
