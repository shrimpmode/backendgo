package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/models"

	"gorm.io/gorm"
)

type CreateUserDTO struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
}

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createUserDTO CreateUserDTO
		fmt.Println("create user")

		if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user := models.User{
			Email:    createUserDTO.Email,
			UserName: createUserDTO.UserName,
		}

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
