package main

import (
	"fmt"
	"time"
)

func tickerRoutine(dur time.Duration) chan int {
	tickerCh := make(chan int, 100)
	go func() {
		ticker := time.NewTicker(dur)
		i := 0
		for {
			select {
			case <-ticker.C:
				fmt.Println("[ticker-routine] tick ", i)
				tickerCh <- i
				i++
				break

			case val := <-tickerCh:
				if val == 0 {
					close(tickerCh)
					return
				}

			}
		}
	}()

	return tickerCh
}
