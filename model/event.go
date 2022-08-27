package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Generalize events, this is just an example
type CommitEvent struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	Project   string             `bson:"project,omitempty"`
	Repo      string             `bson:"repo,omitempty"`
}
