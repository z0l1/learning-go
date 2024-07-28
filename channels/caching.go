package main

import "time"

type Data struct {
	Id        int
	CreatedAt time.Time
}

type DataStore map[string]Data

type DataStoreChan chan DataStore

var ch DataStoreChan

func Init() {
	ch = make(DataStoreChan, 1)
}
