package storage

import (
	"context"
	"devxstats/model"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type store interface {
	AddEvent(model.CommitEvent) error
	GetEvents(projectName string) ([]model.CommitEvent, error)
}

type storeImpl struct {
	db *mongo.Client
}

func InitializeDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	initStore(&storeImpl{db: client})
	if err != nil {
		log.Fatal("could not connect storage", err)
	}
}

var DBStore store

func initStore(store store) {
	DBStore = store
}
