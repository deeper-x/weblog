package web

import (
	"log"
	"net/http"

	"github.com/deeper-x/weblog/messages"
)

// Run the web server
func Run() {
	log.Println(messages.StartServer)

	http.HandleFunc("/save", save)
	http.HandleFunc("/load", load)

	http.ListenAndServe(":8080", nil)
}
