package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/auth/inputs"

	"gorm.io/gorm"
)

func SignUpHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestInput inputs.SignUpInput
		json.NewDecoder(r.Body).Decode(&requestInput)
		fmt.Println("sign up", requestInput)
	}
}
