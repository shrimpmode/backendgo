package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"webserver/auth/inputs"
	"webserver/auth/passwords"
	"webserver/db"
	"webserver/models"

	"github.com/joho/godotenv"
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

type SignUpHandler struct {
	inputReader InputReader[inputs.SignUpInput]
}

func (h *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input, err := h.inputReader.GetInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	success, err := verifyToken(input.Token)
	if !success {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	pass, err := passwords.HashPassword(input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{
		Email:       input.Email,
		Password:    pass,
		Username:    input.Username,
		DisplayName: input.DisplayName,
	}

	err = db.GetDB().Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewSignUpHandler() *SignUpHandler {
	return &SignUpHandler{
		inputReader: &Input[inputs.SignUpInput]{},
	}
}
