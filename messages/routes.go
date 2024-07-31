package messages

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/server/{serverID}/chats/{chatID}/message", CreateMessage(db)).Methods("POST")
}

func CreateMessage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create message")
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Vars: %v", vars)
	}
}
