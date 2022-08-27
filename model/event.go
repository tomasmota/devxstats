package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Commit struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	System    string             `bson:"system,omitempty"` // e.g. Github, Bitbucket
	Group     string             `bson:"group,omitempty"`  // Github org, Bitbucket project
	Repo      string             `bson:"repo,omitempty"`
	User      string             `bson:"user,omitempty"`
}

type Deployment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	System    string             `bson:"system,omitempty"`
	Group     string             `bson:"group,omitempty"`
	Project   string             `bson:"project,omitempty"`
}
