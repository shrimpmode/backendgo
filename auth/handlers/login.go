package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/auth/inputs"
	"webserver/auth/jwt"
	"webserver/auth/passwords"
	"webserver/models"

	"github.com/go-playground/validator/v10"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginInput inputs.LoginInput
		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(&loginInput); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Where("email = ?", loginInput.Email).First(&user).Error; err != nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		if !passwords.CheckPasswordHash(loginInput.Password, user.Password) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token, err := jwt.CreateToken(&user)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		data := inputs.TokenResponse{
			Token: token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}
