package main

import (
	"context"
	"time"

	"github.com/cheynewallace/tabby"
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
	Unit      string             `bson:"type,omitempty"`
	Value     float64            `bson:"value,omitempty"`
}

func mongomain() {
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

	CreateMetrics(metricsCollection)
	AddMeasurement(metricsCollection, measurementCollection)

	t := tabby.New()
	t.AddHeader("METRIC", "TIMESTAMP", "VALUE")
	OutputMeasurements(metricsCollection, measurementCollection, t)
	t.Print()
}

func GetMetric(name string, metricsCollection *mongo.Collection) Metric {
	var metric Metric
	err := metricsCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&metric)
	if err != nil {
		panic(err)
	}
	return metric
}

// hmmm not so practical...
// TODO: find out how to manage collections
func GetMetricMeasurements(name string, metricsCollection *mongo.Collection, measurementCollection *mongo.Collection) []Measurement {
	metric := GetMetric(name, metricsCollection)

	cursor, err := measurementCollection.Find(context.TODO(), bson.M{"metric": metric.ID})
	if err != nil {
		panic(err)
	}

	var measurements []Measurement
	if err = cursor.All(context.TODO(), &measurements); err != nil {
		panic(err)
	}
	return measurements
}

func OutputMeasurements(metricsCollection *mongo.Collection, measurementCollection *mongo.Collection, t *tabby.Tabby) {
	measurements := GetMetricMeasurements("Lead Time", metricsCollection, measurementCollection)
	for _, measurement := range measurements {
		t.AddLine("Lead Time", measurement.Timestamp, time.Duration(measurement.Value))
	}
}

func AddMeasurement(metricsCollection *mongo.Collection, measurementCollection *mongo.Collection) {
	metric := GetMetric("Lead Time", metricsCollection)
	measurement := Measurement{
		Metric:    metric.ID,
		Timestamp: time.Now(),
		Unit:      "duration",
		Value:     float64((time.Hour * 24 * 3)),
	}

	_, err := measurementCollection.InsertOne(context.TODO(), measurement)
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

	_, err := metricsCollection.InsertMany(context.TODO(), metrics)
	if err != nil {
		panic(err)
	}
}
