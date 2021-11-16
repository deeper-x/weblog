package messages

import "fmt"

// StartServer is the debug message when the starting the server
const StartServer = "Starting server..."

// SavingEntry is the debug message when the saving an entry
var SavingEntry = func(sender string) string {
	return fmt.Sprintf("Saving entry - Signature: [%s]", sender)
}
