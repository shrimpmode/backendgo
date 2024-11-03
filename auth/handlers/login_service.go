package handlers

import (
	"errors"
	"webserver/auth/inputs"
	"webserver/auth/jwt"
	"webserver/auth/passwords"
	"webserver/db"
	"webserver/models"

	"gorm.io/gorm"
)

type LoginRepository interface {
	GetUserByEmail(email string) (models.User, error)
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

type LoginService interface {
	GetUserByEmail(email string) (models.User, error)
	GenerateToken(input inputs.LoginInput, user models.User) (string, error)
}

type Service struct {
	repository LoginRepository
}

func (s *Service) GetUserByEmail(email string) (models.User, error) {
	return s.repository.GetUserByEmail(email)
}

func (s *Service) GenerateToken(input inputs.LoginInput, user models.User) (string, error) {
	if !passwords.CheckPasswordHash(input.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.CreateToken(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewLoginService() *Service {
	return &Service{
		repository: &Repository{DB: db.GetDB()},
	}
}
