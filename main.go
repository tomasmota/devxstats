package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Metric struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
}

type Measurement struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Metric    primitive.ObjectID `bson:"metric,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	Value     float64            `bson:"value,omitempty"`
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	metricsCollection := client.Database("devxmetrics").Collection("metrics")
	measurementCollection := client.Database("devxmetrics").Collection("measurements")
	defer metricsCollection.Drop(context.TODO())
	defer measurementCollection.Drop(context.TODO())

	fmt.Println("Creating metrics")
	CreateMetrics(metricsCollection)

	fmt.Println("Search metrics")
	ReadMetrics(metricsCollection)

	fmt.Println("Add Lead Time measurement")
	AddMeasurement(metricsCollection, measurementCollection)

	fmt.Println("Read Measurement")
	// ReadMeasurement(metricsCollection)
}

func AddMeasurement(metricsCollection *mongo.Collection, measurementCollection *mongo.Collection) {
	filter := bson.M{"name": "Lead Time"}

	var metric Metric
	err := metricsCollection.FindOne(context.TODO(), filter).Decode(&metric)
	if err != nil {
		panic(err)
	}

	measurement := Measurement{
		Metric:    metric.ID,
		Timestamp: time.Now(),
		Value:     (time.Hour * 24 * 3).Seconds(),
	}

	_, err = measurementCollection.InsertOne(context.TODO(), measurement)
	if err != nil {
		panic(err)
	}
}

func CreateMetrics(metricsCollection *mongo.Collection) {
	metrics := []interface{}{
		Metric{
			Name:        "Deployment Frequency",
			Description: "Rolling Average number of deployments per week over the past 2 months",
		},
		Metric{
			Name:        "Lead Time",
			Description: "Average duration between commits being pushed and being deployed to production",
		},
		Metric{
			Name:        "Review Time",
			Description: "Average time between a PR being created and being reviewed",
		},
	}

	insertResult, err := metricsCollection.InsertMany(context.TODO(), metrics)
	if err != nil {
		panic(err)
	}

	contactIDs := insertResult.InsertedIDs

	var contactIDs_ []primitive.ObjectID
	for _, id := range contactIDs {
		contactIDs_ = append(contactIDs_, id.(primitive.ObjectID))
	}

	fmt.Printf("Inserted %v %T\n", contactIDs_, contactIDs_)

}

func ReadMetrics(metricsCollection *mongo.Collection) {
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "name", Value: "Lead Time"},
				},
			},
		},
	}

	cursor, err := metricsCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var metrics []Metric
	if err = cursor.All(context.TODO(), &metrics); err != nil {
		panic(err)
	}

	fmt.Println("Matching Metrics:")
	for _, metric := range metrics {
		fmt.Println("  " + metric.Name + ": " + metric.Description)
	}

}
