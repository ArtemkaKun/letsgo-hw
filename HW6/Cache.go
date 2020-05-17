package main

import (
	"errors"
	"fmt"
)

type ID int

var IdCounter ID
var WorkersCache = map[ID]Worker{}

func InsertNewWorker(newWorker interface{}) {
	WorkersCache[IdCounter] = Worker{Id: IdCounter, WorkerInfo: newWorker}
	IdCounter++
}

func GetUserInfoById(id ID) (workerInfo Worker, err error) {
	if data, ok := WorkersCache[id]; ok {
		workerInfo = data
	} else {
		err = errors.New(fmt.Sprintf("Worker with this ID (%v) doesn't exist\n", id))
	}

	return
}

func UpdateWorkerInfo(id ID, newWorkerInfo interface{}) (err error) {
	if _, ok := WorkersCache[id]; ok {
		WorkersCache[id] = Worker{Id: id, WorkerInfo: newWorkerInfo}
	} else {
		err = errors.New(fmt.Sprintf("Worker with this ID (%v) doesn't exist\n", id))
	}

	return
}

func DeleteWorker(id ID) (err error)  {
	if _, ok := WorkersCache[id]; ok {
		delete(WorkersCache, id)
	} else {
		err = errors.New(fmt.Sprintf("Worker with this ID (%v) doesn't exist\n", id))
	}

	return
}
