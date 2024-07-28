package main

import (
	"github.com/google/uuid"
	"math/rand/v2"
	"time"
)

func createRandomData() Data {
	return Data{
		Uuid:      uuid.New(),
		CreatedAt: time.Time{},
	}
}

func getRandomLifespanBetween(max, min int) time.Duration {
	secN := rand.IntN(max-min) + min
	return time.Second * time.Duration(secN)
}
