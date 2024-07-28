package main

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

func formatUuid(uuid uuid.UUID) string {
	first, _, _ := strings.Cut(uuid.String(), "-")
	return first
}

type Data struct {
	Uuid      uuid.UUID
	CreatedAt time.Time
}

func (d Data) String() string {
	return formatUuid(d.Uuid)
}

type DataStore map[uuid.UUID]Data

type DataStoreChan chan DataStore

func CreateDataStore() DataStoreChan {
	ch := make(DataStoreChan, 1)
	ch <- make(DataStore)
	return ch
}

func AddRoutine(ch DataStoreChan, data Data) {
	current := <-ch
	current[data.Uuid] = data
	ch <- current
}

func RemoveRoutine(ch DataStoreChan, uuid uuid.UUID) {
	current := <-ch
	delete(current, uuid)
	ch <- current
}

func GetByIdRoutine(ch DataStoreChan, uuid uuid.UUID) (Data, bool) {
	current := <-ch
	value, ok := current[uuid]
	ch <- current
	return value, ok
}

func GetStateRoutine(ch DataStoreChan) DataStore {
	current := <-ch
	ch <- current
	return current
}
