mongodb web app - web interface for logging events

```golang
	// 1. Create a new database engine
	e := db.NewEngine()
	defer e.Close()

	// 2. Set the context
	cancel := e.SetContext()
	defer cancel()

	// 3. Connect to the database
	e.Connect()

	// 4. Create a new entry
	e.SetCollection("test", "events")
	e.AddEntry("demo xyz")
```