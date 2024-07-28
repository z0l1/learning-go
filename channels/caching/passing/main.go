package main

import (
	"fmt"
	"time"
)

func simulateDataLifecycleRoutine(ch DataStoreChan) {
	data := createRandomData()
	lifespan := getRandomLifespanBetween(10, 2)
	fmt.Println("[created]", data, "for ", lifespan, "s")

	go AddRoutine(ch, data)
	fmt.Println("[launchd]", data, "add")

	time.Sleep(lifespan)
	go RemoveRoutine(ch, data.Uuid)
	fmt.Println("[launchd]", data, "remove")
}

func main() {
	ch := CreateDataStore()

	for range 100 {
		go simulateDataLifecycleRoutine(ch)
	}

	for {
		time.Sleep(time.Millisecond * 500)
		current := GetStateRoutine(ch)
		var ids []string
		for _, data := range current {
			ids = append(ids, data.String())
		}
		fmt.Println("[current] ", ids)

	}

}
