package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/deeper-x/weblog/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewEntry creates a new MS
func NewEntry(sender, text string, ts time.Time) *Entry {
	return &Entry{
		Signature: sender,
		Text:      text,
		TS:        ts,
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

// createClient creates a new client
func (i *Instance) createClient() {
	uri := fmt.Sprintf("mongodb://%s:%s", i.host, i.port)
	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}

	i.client = c
}

// createCollection creates a new collection
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
func (i *Instance) AddEntry(signature, txt string) (bool, error) {
	data := []interface{}{
		NewEntry(signature, txt, time.Now()),
	}

	_, err := i.collection.InsertMany(
		i.ctx,
		data,
	)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

// SaveEntry is the db wrapper to save an entry
func SaveEntry(signature, entry string) (string, error) {
	inst := NewInstance(settings.Host, settings.Port)
	defer inst.Close()

	// Connect to the database
	close, err := inst.Connect(settings.Database, settings.Collection)
	if err != nil {
		log.Println(err)
		return "Connection error", err
	}
	defer close()

	// Create a new entry
	ok, err := inst.AddEntry(signature, entry)
	if err != nil {
		log.Println(err)
		return "DB save error", err
	}

	output := strconv.FormatBool(ok)
	return output, nil
}
