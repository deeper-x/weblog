package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deeper-x/weblog/db"
	"github.com/deeper-x/weblog/messages"
	"github.com/deeper-x/weblog/wauth"
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
	message := req.URL.Query().Get("message")

	isAuth, err := wauth.IsAllowed(signature)
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(messages.AuthError))
		return
	}

	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(messages.AuthDenied))

		return
	}

	if !checkSaveParams(signature, message) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(messages.MissingParamsErr))

		return
	}

	// Create a new database engine
	smessage := messages.SaveMsg(signature)
	log.Println(smessage)

	res, err := db.SaveEntry(signature, message)
	if err != nil {
		log.Println(messages.SavingErr, err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(messages.SavingErr))

		return
	}

	log.Println(messages.Saved)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func load(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/json: charset=utf-8")

	// reading GET parameters signature
	signature := req.URL.Query().Get("signature")

	isAuth, err := wauth.IsAllowed(signature)
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(messages.AuthError))

		return
	}

	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(messages.AuthDenied))

		return
	}

	if !checkLoadParam(signature) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(messages.MissingParamsErr))

		return
	}

	data, err := db.GetEntries(signature)
	if err != nil {
		log.Println(messages.Loaded, err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(messages.Loaded))

		return
	}

	jData, err := json.Marshal(data)
	if err != nil {
		log.Println(messages.Loaded, err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(messages.Loaded))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

// checkSaveParams checks if the GET parameters is valid
func checkSaveParams(sender, entry string) bool {
	if len(sender) == 0 || len(entry) == 0 {
		return false
	}

	return true
}

// checkLoadParam checks if the GET parameters is valid
func checkLoadParam(sender string) bool {
	return len(sender) != 0
}
