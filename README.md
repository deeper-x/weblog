mongodb web app - web interface for logging events

```golang
	// Create a new database engine
	inst := db.NewInstance("localhost", "27017")
	defer inst.Close()

	// Connect to the database, picking a collection
	close, err := inst.Connect("test", "events")
	if err != nil {
		log.Panic(err)
	}
	defer close()

	// Create a new entry
	inst.AddEntry("System XYZ - Failure detected")
```