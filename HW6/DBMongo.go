package main

import (
	"fmt"
	_ "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"log"
)

type MongoObj map[string]interface{}

const(
	DBName string = "DBWork"
	CollectionName string = "Workers"
)

var mongoConnection = func() (session *mgo.Session) {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Connected to Mongo!")
	return
}()

var collectionWorkers = func() (data *mgo.Collection) {
	data = mongoConnection.DB(DBName).C(CollectionName)

	if names, _ := data.Database.CollectionNames(); names == nil {
		log.Panicf("Database %v doesn't exist!\n", DBName)
	}

	log.Printf("Connected to %v:%v\n", DBName, CollectionName)
	return
}()

func InsertWorkerInDB(newWorker Worker) {
	err := collectionWorkers.Insert(newWorker)
	if err != nil {
		fmt.Println(err)
	}
}

func GetWorkerDataFromDB(id ID) (result Worker) {
	err := collectionWorkers.Find(MongoObj{"_id": id}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func UpdateWorkerDataDB(id ID, newData Worker) {
	err := collectionWorkers.Update(MongoObj{"_id": id}, newData)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteWorkerDataFromDB(id ID) {
	err := collectionWorkers.Remove(MongoObj{"_id": id})
	if err != nil {
		fmt.Println(err)
	}
}
