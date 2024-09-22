package jwt

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"webserver/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func CreateToken(user *models.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func ParseToken(tokenString string) (jwt.MapClaims, bool) {
	err := godotenv.Load()
	if err != nil {
		return nil, false
	}
	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

func GetTokenFromRequest(r *http.Request) (string, bool) {
	authorization := r.Header.Get("authorization")

	strs := strings.Split(authorization, "Bearer ")
	if len(strs) == 2 {
		return strs[1], true
	}
	return "", false
}

type JWTAuthenticator struct {
	db *gorm.DB
}

func (authenticator *JWTAuthenticator) GetAuthenticatedUser(r *http.Request) (*models.User, bool) {
	user := models.User{}
	tokenString, ok := GetTokenFromRequest(r)
	if !ok {
		return nil, false
	}
	claims, ok := ParseToken(tokenString)
	if !ok {
		return nil, false
	}
	email := claims["email"]
	result := authenticator.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, false
	}

	return &user, ok
}

func NewJWTAuthenticator(db *gorm.DB) *JWTAuthenticator {
	return &JWTAuthenticator{db}
}
