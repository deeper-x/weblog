package main

import (
	"github.com/deeper-x/weblog/db"
)

func main() {
	// Create a new database engine
	e := db.NewEngine()
	defer e.Close()

	// Set the context
	cancel := e.SetContext()
	defer cancel()

	// Connect to the database
	e.Connect()

	// Create a new entry
	e.SetCollection("test", "events")
	e.AddEntry("demo xyz")
}
