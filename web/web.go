package web

import (
	"log"
	"net/http"

	"github.com/deeper-x/weblog/db"
	"github.com/deeper-x/weblog/messages"
)

func test(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html: charset=utf-8")
	val := req.URL.Query().Get("demo")

	if len(val) > 0 {
		log.Println("val is", val)
	}

	w.Write([]byte(`<p>hello</p>`))
}

func save(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html: charset=utf-8")

	// reading GET parameters
	signature := req.URL.Query().Get("signature")
	entry := req.URL.Query().Get("entry")

	if !checkParams(signature, entry) {
		w.Write([]byte(messages.MissingParamsErr))
		return
	}

	// Create a new database engine
	smessage := messages.SavingEntry(signature)
	log.Println(smessage)

	res, err := db.SaveEntry(signature, entry)
	if err != nil {
		log.Println(messages.SavingErr, err)
		w.Write([]byte(messages.SavingErr))
		return
	}

	w.Write([]byte(res))
}

func checkParams(sender, entry string) bool {
	if len(sender) == 0 || len(entry) == 0 {
		return false
	}

	return true
}

// Run the web server
func Run() {
	log.Println(messages.StartServer)

	http.HandleFunc("/save", save)
	http.ListenAndServe(":8080", nil)
}
