package storage

import (
	"context"
	"devxstats/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection = "repos"

func (storeImpl *storeImpl) AddRepos(ctx context.Context, repos []*model.Repository) error {
	// iRepos := make([]interface{}, len(repos))

	// for i := range repos {
	// 	iRepos[i] = repos[i]
	// }
	collection := storeImpl.db.Database("devxstats").Collection(collection)
	fmt.Printf("Inserting %v repos\n", len(repos))

	var ops []mongo.WriteModel
	for _, r := range repos {
		var ir interface{} = r
		op := mongo.NewInsertOneModel()
		op.SetDocument(ir)
		// op.SetFilter(bson.M{"System": r.System, "Group": r.Group, "Name": r.Name})
		// op.SetUpdate(ir)
		// op.SetUpsert(true)
		ops = append(ops, op)
	}

	opts := options.BulkWriteOptions{}
	opts.SetOrdered(true)
	res, err := collection.BulkWrite(ctx, ops, &opts)

	// res, err := reposCollection.UpdateMany(ctx, bson.M{}, iRepos, &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return fmt.Errorf("error bulk inserting repos: %v", err)
	}
	fmt.Printf("%v repos stored", res.ModifiedCount)
	fmt.Printf("%v repos already matched", res.MatchedCount)
	// _, err := reposCollection.InsertMany(ctx, iRepos, &options.InsertManyOptions{})
	return nil
}

func (storeImpl *storeImpl) GetRepos(group string) ([]model.Repository, error) {
	reposCollection := storeImpl.db.Database("devxstats").Collection(collection)
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
