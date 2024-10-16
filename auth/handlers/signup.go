package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"webserver/auth/inputs"
	"webserver/auth/passwords"
	"webserver/models"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func verifyToken(token string) (bool, error) {
	err := godotenv.Load()
	if err != nil {
		return false, err
	}

	env := os.Getenv("GO_ENV")
	if env == "development" {
		return true, nil
	}

	recaptchaUrl := os.Getenv("CAPTCHA_URL")
	secret := os.Getenv("RE_SECRET_KEY")

	form := url.Values{"response": {token}, "secret": {secret}}

	resp, err := http.PostForm(recaptchaUrl, form)
	if err != nil {
		return false, err
	}

	var captchaResponse inputs.RecaptchaResponse

	err = json.NewDecoder(resp.Body).Decode(&captchaResponse)
	if err != nil {
		return false, err
	}

	return captchaResponse.Success, err
}

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

		success, err := verifyToken(requestInput.Token)
		if !success {
			http.Error(w, err.Error(), http.StatusForbidden)
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
	}
}
