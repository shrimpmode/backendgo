package messages

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/server/{serverID}/chats/{chatID}/message", CreateMessage(db)).Methods("POST")
}

func CreateMessage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Create message")
		vars := mux.Vars(r)
		log.Printf("Vars: %v", vars)
	}
}
