package storage

import (
	"context"
	"devxstats/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (storeImpl *storeImpl) AddEvent(event model.CommitEvent) error {
	eventsCollection := storeImpl.db.Database("devxstats").Collection("events")

	insertResult, err := eventsCollection.InsertOne(context.TODO(), event)
	if err != nil {
		return err
	}
	log.Println(insertResult) //TODO: remove this
	return nil
}

func (storeImpl *storeImpl) GetEvents(projectName string) (model.CommitEvent, error) {
	eventsCollection := storeImpl.db.Database("devxstats").Collection("events")

	var event model.CommitEvent
	err := eventsCollection.FindOne(context.TODO(), bson.M{"project": projectName}).Decode(&event)
	if err != nil {
		panic(err)
	}
	return event, nil
}
