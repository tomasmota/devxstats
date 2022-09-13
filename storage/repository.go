package storage

import (
	"context"
	"devxstats/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (storeImpl *storeImpl) AddRepos(ctx context.Context, repos []*model.Repository) error {
	iRepos := make([]interface{}, len(repos))

	for i := range repos {
		iRepos[i] = repos[i]
	}
	reposCollection := storeImpl.db.Database("devxstats").Collection("repositories")
	fmt.Printf("Inserting %v repos\n", len(repos))
	_, err := reposCollection.InsertMany(ctx, iRepos, &options.InsertManyOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) GetRepos(group string) ([]model.Repository, error) {
	reposCollection := storeImpl.db.Database("devxstats").Collection("repositories")

	var repos []model.Repository
	cursor, err := reposCollection.Find(context.TODO(), bson.M{"group": group})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &repos); err != nil {
		return nil, err
	}

	return repos, nil
}
