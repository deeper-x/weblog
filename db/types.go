package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Instance is the interface for the database
type Instance struct {
	host       string
	port       string
	client     *mongo.Client
	ctx        context.Context
	collection *mongo.Collection
	Cancel     context.CancelFunc
}

// Entry is a MongoDB object
type Entry struct {
	Signature string
	TS        time.Time
	Message   string
}
