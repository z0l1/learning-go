package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func runBoolIndicator() {
	n := 100

	tickerCh := tickerRoutine(time.Second * 1)
	for range n {
		fmt.Println("blocking until reading value")
		tick, chOk := <-tickerCh
		fmt.Println("tick: ", tick, "; ", chOk)

		if !chOk {
			fmt.Println("channel got closed")
			break
		}

		rnd := rand.IntN(10)
		fmt.Println("randomed", rnd)
		if rnd > 5 {
			tickerCh <- 0
		}
	}
}
