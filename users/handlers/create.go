package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/auth"
	"webserver/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateUserDTO struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validate := validator.New(validator.WithRequiredStructEnabled())
		var createUserDTO CreateUserDTO

		if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&createUserDTO); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		password, _ := auth.HashPassword(createUserDTO.Password)

		user := models.User{
			Email:    createUserDTO.Email,
			UserName: createUserDTO.UserName,
			Password: password,
		}

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
