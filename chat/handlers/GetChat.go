package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetChat(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chats := []models.Chat{}
		vars := mux.Vars(r)

		where := map[string]interface{}{
			"server_id": vars["serverId"],
		}
		db.Where(where).Find(&chats)
		fmt.Println("Get chats", vars, chats)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chats)
	}
}
