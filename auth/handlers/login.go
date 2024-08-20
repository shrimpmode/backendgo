package handlers

import (
	"encoding/json"
	"fmt"
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
		validate := validator.New(validator.WithRequiredStructEnabled())

		var loginInput inputs.LoginInput
		var user models.User

		json.NewDecoder(r.Body).Decode(&loginInput)

		if err := validate.Struct(&loginInput); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.Where("email = ?", loginInput.Email).First(&user)

		isValid := passwords.CheckPasswordHash(loginInput.Password, user.Password)

		if !isValid {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		session, _ := store.Get(r, "cookie-name")

		session.Values["authenticated"] = true
		session.Save(r, w)

		token, err := jwt.CreateToken(&user)
		fmt.Println(token, err)
	}
}
