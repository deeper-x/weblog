package settings

import "fmt"

// Host is the hostname of the database
const Host = "localhost"

// Port is the port of the database
const Port = "27017"

// Database is the name of the database
const Database = "test"

// Collection is the username of the database
const Collection = "events"

// RootDir is the root directory of the project
const RootDir = "/home/deeper-x/go/src/github.com/deeper-x/weblog"

// GetAuthFile is the path to the file containing the authentication information
var GetAuthFile = func() string {
	return fmt.Sprintf("%s/wauth/whitelist.json", RootDir)
}
