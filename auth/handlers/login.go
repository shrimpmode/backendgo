package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/auth/dtos"
	"webserver/auth/passwords"
	"webserver/models"

	"github.com/go-playground/validator/v10"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validate := validator.New(validator.WithRequiredStructEnabled())
		var loginDTO dtos.LoginDTO
		var user models.User

		json.NewDecoder(r.Body).Decode(&loginDTO)

		if err := validate.Struct(&loginDTO); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.Where("email = ?", loginDTO.Email).First(&user)

		isValid := passwords.CheckPasswordHash("test", user.Password)

		if !isValid {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		session, _ := store.Get(r, "cookie-name")

		session.Values["authenticated"] = true
		session.Save(r, w)
	}
}
