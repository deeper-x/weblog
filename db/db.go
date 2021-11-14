package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Engine is the interface for the database
type Engine struct {
	client     *mongo.Client
	ctx        context.Context
	collection *mongo.Collection
}

// Entry is a MongoDB object
type Entry struct {
	TS   time.Time
	Text string
}

// NewEntry creates a new MS
func NewEntry(text string, ts time.Time) *Entry {
	return &Entry{
		Text: text,
		TS:   ts,
	}
}

// NewEngine creates a new database engine
func NewEngine() *Engine {
	return &Engine{}
}

// SetContext sets the context
func (e *Engine) SetContext() context.CancelFunc {
	var cancel context.CancelFunc

	e.ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	return cancel
}

// Connect to the MongoDB database
func (e *Engine) Connect() (*mongo.Client, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Panic(err)
	}

	e.client = c

	err = c.Connect(e.ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return c, nil
}

// SetCollection returns a collection
func (e *Engine) SetCollection(dbname, collection string) {
	e.collection = e.client.Database(dbname).Collection(collection)
}

// Close the database connection, to be deferred
func (e *Engine) Close() {
	e.client.Disconnect(e.ctx)
}

// AddEntry adds an entry to the database
func (e *Engine) AddEntry(msg string) error {
	data := []interface{}{
		NewEntry(msg, time.Now()),
	}

	res, err := e.collection.InsertMany(
		e.ctx,
		data,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(res.InsertedIDs...)
	return nil
}
