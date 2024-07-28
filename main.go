package main

import (
	"fmt"
	"net/http"
	"webserver/db"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/server/{serverID}/chats/{chatID}/message", CreateMessage).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create message")
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Vars: %v", vars)
}
