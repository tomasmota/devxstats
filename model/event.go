package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Build struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StartTime time.Time          `bson:"starttime,omitempty"`
	EndTime   time.Time          `bson:"endtime,omitempty"`
	Team      string             `bson:"team,omitempty"`
	System    string             `bson:"system,omitempty"` // e.g. Tekton, Teamcity
	Group     string             `bson:"group,omitempty"`  // tekton namespace, teamcity high-level project
	Project   string             `bson:"repo,omitempty"`   // tekton pipeline, teamcity project
	User      string             `bson:"user,omitempty"`
	Succeeded bool               `bson:"succeeded,omitempty"`
}

type PullRequest struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OpenTime  time.Time          `bson:"opentime,omitempty"`
	CloseTime time.Time          `bson:"closetime,omitempty"`
	Status    time.Time          `bson:"status,omitempty"`
	Team      string             `bson:"team,omitempty"`
	System    string             `bson:"system,omitempty"` // e.g. Github, Bitbucket
	Group     string             `bson:"group,omitempty"`  // Github org, Bitbucket project
	Repo      string             `bson:"repo,omitempty"`
	User      string             `bson:"user,omitempty"`
}

type Commit struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	Team      string             `bson:"team,omitempty"`
	System    string             `bson:"system,omitempty"` // e.g. Github, Bitbucket
	Group     string             `bson:"group,omitempty"`  // Github org, Bitbucket project
	Repo      string             `bson:"repo,omitempty"`
	User      string             `bson:"user,omitempty"`
}

type Deployment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StartTime time.Time          `bson:"starttime,omitempty"`
	EndTime   time.Time          `bson:"endtime,omitempty"`
	Team      string             `bson:"team,omitempty"`
	System    string             `bson:"system,omitempty"`
	Group     string             `bson:"group,omitempty"`
	Project   string             `bson:"project,omitempty"`
	Succeeded bool               `bson:"succeeded,omitempty"`
}
