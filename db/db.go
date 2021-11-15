package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

// createCtx sets the context
func (i *Instance) createCtx() {
	i.ctx, i.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
}

func (i *Instance) createClient() {
	uri := fmt.Sprintf("mongodb://%s:%s", i.host, i.port)
	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}

	i.client = c
}

func (i *Instance) createCollection(dbname, collname string) {
	i.collection = i.client.Database(dbname).Collection(collname)
}

// Connect to the MongoDB database
func (i *Instance) Connect(dbname, collname string) (context.CancelFunc, error) {
	// start client, define a collection and set the context
	i.createClient()
	i.createCollection(dbname, collname)
	i.createCtx()

	// then connect
	err := i.client.Connect(i.ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// if successuful connected return the context
	return i.Cancel, nil
}

// Close the database connection, to be deferred
func (i *Instance) Close() {
	i.client.Disconnect(i.ctx)
	i.client = nil
}

// AddEntry adds an entry to the database
func (i *Instance) AddEntry(msg string) (*mongo.InsertManyResult, error) {
	data := []interface{}{
		NewEntry(msg, time.Now()),
	}

	res, err := i.collection.InsertMany(
		i.ctx,
		data,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}
