package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// createCtx sets the context
func (e *Instance) createCtx() context.CancelFunc {
	e.ctx, e.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	return e.Cancel
}

func (e *Instance) createClient() {
	uri := fmt.Sprintf("mongodb://%s:%s", e.host, e.port)
	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}

	e.client = c
}

func (e *Instance) createCollection(dbname, collname string) {
	e.collection = e.client.Database(dbname).Collection(collname)
}

// NewEntry creates a new MS
func NewEntry(text string, ts time.Time) *Entry {
	return &Entry{
		Text: text,
		TS:   ts,
	}
}

// NewInstance creates a new database engine
func NewInstance(host, port string) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

// Connect to the MongoDB database
func (e *Instance) Connect(dbname, collname string) (context.CancelFunc, error) {
	// start client, define a collection and set the context
	e.createClient()
	e.createCollection(dbname, collname)
	e.createCtx()

	// then connect
	err := e.client.Connect(e.ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// if successuful connected return the context
	return e.Cancel, nil
}

// Close the database connection, to be deferred
func (e *Instance) Close() {
	e.client.Disconnect(e.ctx)
}

// AddEntry adds an entry to the database
func (e *Instance) AddEntry(msg string) (*mongo.InsertManyResult, error) {

	data := []interface{}{
		NewEntry(msg, time.Now()),
	}

	res, err := e.collection.InsertMany(
		e.ctx,
		data,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}
