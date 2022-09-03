package storage

import (
	"context"
	"devxstats/model"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type store interface {
	AddCommits([]interface{}) error
	GetCommits(projectName string) ([]model.Commit, error)
}

type storeImpl struct {
	db *mongo.Client
}

func InitializeDB(dbUri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(fmt.Errorf("an error occured while creating database connection: %v", err))
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Errorf("an error occured while pinging database: %v", err))
	}

	initStore(&storeImpl{db: client})
}

var DBStore store

func initStore(store store) {
	DBStore = store
}
