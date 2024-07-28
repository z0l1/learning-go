package linkedlist

import (
	"fmt"
	"time"
)

func RunTest() {

	ll := New[int]()

	fmt.Println(ll.Remove(0))
	fmt.Println(ll.Remove(1000))

	ll.Add(1)
	fmt.Println(ll.Remove(0))

	ll.Add(2)
	ll.Add(3)
	fmt.Println(ll.Remove(1))

	//fmt.Println(ll.Get(-1))
	fmt.Println(ll.Get(0))

	for i := 4; i < 100000; i++ {
		ll.Add(i)

		if i%5 == 0 {
			fmt.Println(ll.Remove(uint(i) - 1))
		}
	}

	start := time.Now()
	for val := range ll.ChanIter(false) {
		v := val
		v = v + v
	}
	fmt.Println("iterating", ll.Len(), "items took:", time.Since(start))

	//for i := range ll {
	//
	//}
}
