package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/auth/inputs"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SignUpHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestInput inputs.SignUpInput
		err := json.NewDecoder(r.Body).Decode(&requestInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		validate := validator.New()
		err = validate.Struct(&requestInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		result := db.Create(&requestInput)
		if result.Error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
