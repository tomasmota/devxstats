package model

import (
	"time"
)

// System info is pulled from, e.g. Github, ArgoCD, Tekton, Jira
type System struct {
	ID   int
	Name string
	Type string
}

// High level grouping of resources, might contain repositories or projects within
type Group struct {
	ID       int
	SystemID int
	Name     string
}

// TODO: add this to schema once implementation is started
type Project struct {
	ID      int
	GroupID int
	Name    string
}

type Repo struct {
	ID      int
	GroupID int
	ScmID   int
	Name    string
}

type PullRequest struct {
	ID        int
	RepoID    int
	Number    int
	SourceRef string
	TargetRef string
	Closed    bool
	Merged    bool
	CreatedAd time.Time
	ClosedAt  time.Time
}

type Review struct {
	ID        int
	PrID      int
	CreatedAt time.Time
}

// TODO: add this to schema once implementation is started
type Deployment struct {
	ID        int
	ProjectID int
	StartTime time.Time
	EndTime   time.Time
	System    string
	Project   string
	Succeeded bool
}
