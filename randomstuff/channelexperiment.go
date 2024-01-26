package main

import (
	"fmt"
	"time"
)

func packer(ch chan<- int) {
	for {
		a := 1
		ch <- a
		fmt.Println("packed ", a)
		time.Sleep(time.Millisecond * 500)
	}
}

func picker(ch <-chan int) {
	start := time.Now()
	for {
		fmt.Println("waiting to pick")
		a := <-ch
		fmt.Println("PICKED: ", a)
		time.Sleep(time.Second * 2)
		if time.Since(start) > time.Second*5 {
			fmt.Println("tired of picking")
			return
		}
	}
}
