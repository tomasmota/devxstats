package storage

import (
	"context"
	"devxstats/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (storeImpl *storeImpl) AddEvent(event model.CommitEvent) error {
	eventsCollection := storeImpl.db.Database("devxstats").Collection("events")

	_, err := eventsCollection.InsertOne(context.TODO(), event)
	if err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) GetEvents(projectName string) ([]model.CommitEvent, error) {
	eventsCollection := storeImpl.db.Database("devxstats").Collection("events")

	var events []model.CommitEvent
	cursor, err := eventsCollection.Find(context.TODO(), bson.M{"project": projectName})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &events); err != nil {
		return nil, err
	}

	return events, nil
}
