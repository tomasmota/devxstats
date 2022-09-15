package storage

import (
	"context"
	"devxstats/internal/config"
	"devxstats/model"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type store interface {
	AddCommits([]interface{}) error
	GetCommits(string) ([]model.Commit, error)
	AddRepos(context.Context, []*model.Repository) error
	GetRepos(string) ([]model.Repository, error)
}

type storeImpl struct {
	db *mongo.Client
}

func InitializeDB(ctx context.Context, c *config.DbConfig) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%v:%v", c.Host, c.Port)
	fmt.Println("connecting to database at: ", uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
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
