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

Result:
```mongo
{ "_id" : ObjectId("61926a844e8ea1f1871aedd9"), "ts" : ISODate("2021-11-15T14:11:16.078Z"), "text" : "System XYZ - Failure detected" }

```