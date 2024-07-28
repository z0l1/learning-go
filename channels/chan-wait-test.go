package main

import (
	"fmt"
	"time"
)

func runChanWaitTest() {
	numCh := tickerRoutine(time.Millisecond * 500)

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("[rtn1] waiting for value")
		val, ok := <-numCh
		fmt.Println("[rtn1] value:", val, ok)
	}()

	go func() {
		fmt.Println("[rtn2] waiting for value")
		val, ok := <-numCh
		fmt.Println("[rtn2] value:", val, ok)
	}()

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("[rtn3] waiting for value")
		val, ok := <-numCh
		fmt.Println("[rtn3] value:", val, ok)
	}()

	time.Sleep(time.Minute)
	//fmt.Println("[main] waiting for value")
	//val, ok := <-numCh
	//fmt.Println("[main] value:", val, ok)
}
