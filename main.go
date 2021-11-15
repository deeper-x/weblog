package main

import (
	"log"

	"github.com/deeper-x/weblog/db"
)

func main() {
	// Create a new database engine
	inst := db.NewInstance("localhost", "27017")
	defer inst.Close()

	// Connect to the database
	close, err := inst.Connect("test", "events")
	if err != nil {
		log.Panic(err)
	}
	defer close()

	// Create a new entry
	inst.AddEntry("System XYZ - Failure detected")
}
