package storage

import (
	"context"
	"devxstats/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	reposCollection   = "repos"
	commitsCollection = "commits"
)

var upsert = true

func makeInsert(i interface{}) interface{} {
	return bson.M{
		"$setOnInsert": i,
	}
}

func (storeImpl *storeImpl) AddRepos(ctx context.Context, repos []*model.Repository) error {
	collection := storeImpl.db.Database("devxstats").Collection(reposCollection)
	fmt.Printf("Inserting %v repos\n", len(repos))

	for _, r := range repos {

		r.ID = primitive.NewObjectID()
		filter := bson.M{
			"System": r.System,
			"Group":  r.Group,
			"Name":   r.Name,
		}

		_, err := collection.UpdateOne(ctx, filter, makeInsert(r), &options.UpdateOptions{Upsert: &upsert})
		if err != nil {
			return fmt.Errorf("error inserting repo %v: %w", r.Name, err)
		}
	}

	// print count of upserts here

	return nil
}

func (storeImpl *storeImpl) GetRepos(group string) ([]model.Repository, error) {
	reposCollection := storeImpl.db.Database("devxstats").Collection(reposCollection)
	fmt.Printf("fetching repos in group: %v\n", group)

	var repos []model.Repository
	cursor, err := reposCollection.Find(context.TODO(), bson.M{})
	// cursor, err := reposCollection.Find(context.TODO(), bson.M{"group": group})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func (storeImpl *storeImpl) AddCommits(commits []interface{}) error {
	commitsCollection := storeImpl.db.Database("devxstats").Collection(commitsCollection)

	_, err := commitsCollection.InsertMany(context.TODO(), commits, &options.InsertManyOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) GetCommits(group string) ([]model.Commit, error) {
	commitsCollection := storeImpl.db.Database("devxstats").Collection(commitsCollection)

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
