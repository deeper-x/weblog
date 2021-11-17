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

	err := http.ListenAndServeTLS(":443", "tls/server.crt", "tls/server.key", nil)
	if err != nil {
		log.Panic(err)
	}

}
