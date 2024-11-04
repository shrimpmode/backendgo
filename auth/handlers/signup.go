package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"webserver/app"
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
	inputReader   app.InputReader[inputs.SignUpInput]
	signUpService SignUpService
}

type SignUpService interface {
	VerifyToken(token string) (bool, error)
	HashPassword(password string) (string, error)
	CreateUser(user models.User) error
}

type signUpService struct{}

func (s *signUpService) VerifyToken(token string) (bool, error) {
	return verifyToken(token)
}

func (s *signUpService) HashPassword(password string) (string, error) {
	return passwords.HashPassword(password)
}

func (s *signUpService) CreateUser(user models.User) error {
	return db.GetDB().Create(&user).Error
}

func (h *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input, err := h.inputReader.GetInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	success, err := h.signUpService.VerifyToken(input.Token)
	if !success {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	pass, err := h.signUpService.HashPassword(input.Password)
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

	err = h.signUpService.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewSignUpHandler() *SignUpHandler {
	return &SignUpHandler{
		inputReader:   &app.Input[inputs.SignUpInput]{},
		signUpService: &signUpService{},
	}
}
