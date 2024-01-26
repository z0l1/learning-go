package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 2)
	//go packer(ch)
	go picker(ch)

	for i := 0; ; i++ {
		fmt.Println("packing ")
		ch <- i
	}

	time.Sleep(time.Minute * 1)
}
