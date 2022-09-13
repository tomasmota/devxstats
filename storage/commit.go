package storage

import (
	"context"
	"devxstats/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (storeImpl *storeImpl) AddCommits(commits []interface{}) error {
	commitsCollection := storeImpl.db.Database("devxstats").Collection("commits")

	_, err := commitsCollection.InsertMany(context.TODO(), commits, &options.InsertManyOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) GetCommits(group string) ([]model.Commit, error) {
	commitsCollection := storeImpl.db.Database("devxstats").Collection("commits")

	var commits []model.Commit
	cursor, err := commitsCollection.Find(context.TODO(), bson.M{"group": group})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &commits); err != nil {
		return nil, err
	}

	return commits, nil
}
