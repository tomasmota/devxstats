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

func (storeImpl *storeImpl) AddRepos(ctx context.Context, repos []*model.Repository) error {
	iRepos := make([]interface{}, len(repos))

	for i := range repos {
		repos[i].ID = primitive.NewObjectID()
		iRepos[i] = repos[i]
	}
	collection := storeImpl.db.Database("devxstats").Collection(reposCollection)
	fmt.Printf("Inserting %v repos\n", len(repos))

	// _, err := collection.InsertOne(ctx, iRepos[1], &options.InsertOneOptions{})
	res, err := collection.InsertMany(ctx, iRepos, &options.InsertManyOptions{})
	if err != nil {

		fmt.Println(err)
		// return fmt.Errorf("error inserting repos: %v", err)
	}
	fmt.Printf("inserted %v repos into db\n", len(res.InsertedIDs))
	return nil
}

func (storeImpl *storeImpl) GetRepos(group string) ([]model.Repository, error) {
	reposCollection := storeImpl.db.Database("devxstats").Collection(reposCollection)
	fmt.Printf("fetching repos in group: %v\n", group)

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
