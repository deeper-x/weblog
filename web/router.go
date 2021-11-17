package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deeper-x/weblog/messages"
	"github.com/deeper-x/weblog/settings"
)

// Run the web server
func Run() {
	log.Println(messages.StartServer)

	http.HandleFunc("/save", save)
	http.HandleFunc("/load", load)

	crt := fmt.Sprintf("%s/tls/%s", settings.RootDir, "server.crt")
	key := fmt.Sprintf("%s/tls/%s", settings.RootDir, "server.key")

	err := http.ListenAndServeTLS(":443", crt, key, nil)
	if err != nil {
		log.Panic(err)
	}

}
