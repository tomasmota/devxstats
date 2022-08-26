package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)

	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	metricsCollection := client.Database("devxmetrics").Collection("metrics")

	fmt.Println("Creating metrics")
	CreateMetrics(metricsCollection)
	fmt.Println("Search metrics")
	ReadMetrics(metricsCollection)
	fmt.Println("Delete metrics")
	DeleteMetrics(metricsCollection)
}

func CreateMetrics(metricsCollection *mongo.Collection) {
	users := []interface{}{
		bson.D{{Key: "name", Value: "Deployment Frequency"}, {Key: "description", Value: "Rolling Average number of deployments per week over the past 2 months"}},
		bson.D{{Key: "name", Value: "Lead Time"}, {Key: "description", Value: "Average duration between commits being pushed and being deployed to production"}},
		bson.D{{Key: "name", Value: "Review Time"}, {Key: "description", Value: "Average time between a PR being created and being reviewed"}},
	}

	_, err := metricsCollection.InsertMany(context.TODO(), users)
	if err != nil {
		panic(err)
	}
}

func ReadMetrics(metricsCollection *mongo.Collection) {
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "name", Value: bson.D{{Key: "$regex", Value: "Time"}}},
				},
			},
		},
	}

	cursor, err := metricsCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}

}

func DeleteMetrics(metricsCollection *mongo.Collection) {
	metricsCollection.DeleteMany(context.TODO(), bson.D{})
}
