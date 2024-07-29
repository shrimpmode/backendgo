package messages

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DefineRoutes(r *mux.Router) {
	r.HandleFunc("/server/{serverID}/chats/{chatID}/message", CreateMessage).Methods("POST")
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create message")
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Vars: %v", vars)
}
