package storage

import (
	"context"
	"devxstats/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (storeImpl *storeImpl) AddCommit(event model.Commit) error {
	eventsCollection := storeImpl.db.Database("devxstats").Collection("events")

	_, err := eventsCollection.InsertOne(context.TODO(), event)
	if err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) GetCommits(group string) ([]model.Commit, error) {
	eventsCollection := storeImpl.db.Database("devxstats").Collection("events")

	var events []model.Commit
	cursor, err := eventsCollection.Find(context.TODO(), bson.M{"group": group})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &events); err != nil {
		return nil, err
	}

	return events, nil
}
