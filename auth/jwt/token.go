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

func GetTokenFromRequest(r *http.Request) string {
	authorization := r.Header.Get("authorization")

	strs := strings.Split(authorization, "Bearer ")
	if len(strs) == 2 {
		return strs[1]
	}
	return ""
}

func GetAuthenticatedUser(db *gorm.DB, r *http.Request) (*models.User, bool) {
	user := models.User{}
	tokenString := GetTokenFromRequest(r)
	claims, ok := ParseToken(tokenString)
	email := claims["email"]
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, false
	}

	return &user, ok
}
