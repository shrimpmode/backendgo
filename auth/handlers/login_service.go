package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"webserver/auth/inputs"
	"webserver/auth/jwt"
	"webserver/auth/passwords"
	"webserver/db"
	"webserver/models"

	"github.com/go-playground/validator/v10"
)

type LoginServiceInterface interface {
	GetUser(email string) (models.User, error)
	GetInput(r *http.Request) (inputs.LoginInput, error)
	GenerateToken(input inputs.LoginInput, user models.User) (string, error)
}

type LoginService struct {
}

func (s *LoginService) GetUser(email string) (models.User, error) {
	var user models.User
	if err := db.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *LoginService) GetInput(r *http.Request) (inputs.LoginInput, error) {
	var loginInput inputs.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
		return loginInput, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&loginInput); err != nil {
		return loginInput, err
	}

	return loginInput, nil
}

func (s *LoginService) GenerateToken(input inputs.LoginInput, user models.User) (string, error) {
	if !passwords.CheckPasswordHash(input.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.CreateToken(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}
