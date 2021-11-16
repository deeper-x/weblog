package messages

import "fmt"

// StartServer is the debug message when the starting the server
const StartServer = "Starting server..."

// SaveMsg is the debug message when the saving an entry
var SaveMsg = func(sender string) string {
	return fmt.Sprintf("Saving entry - Signature: [%s]", sender)
}

// Saved is the debug message when the entry is saved
var Saved = "Entry saved succesfully ✔"

// Loaded is the debug message when the entry is loaded
var Loaded = "Entries loaded succesfully ✔"
