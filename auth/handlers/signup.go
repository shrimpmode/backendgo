package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/auth/inputs"
	"webserver/auth/passwords"
	"webserver/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SignUpHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestInput inputs.SignUpInput
		err := json.NewDecoder(r.Body).Decode(&requestInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(&requestInput)
		if err != nil {
			fmt.Println("Failed signup input validation")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pass, err := passwords.HashPassword(requestInput.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user := models.User{
			Email:       requestInput.Email,
			Password:    pass,
			UserName:    requestInput.UserName,
			DisplayName: requestInput.DisplayName,
		}

		result := db.Create(&user)
		if result.Error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("Created user %v", user)
	}
}
