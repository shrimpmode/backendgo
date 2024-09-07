package middleware

import (
	"log"
	"net/http"
	"strings"
	"webserver/auth/jwt"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func IsAuthenticated(db *gorm.DB, store *sessions.CookieStore) Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "cookie-name")
			if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			hf(w, r)
		}
	}
}

func JwtAuthenticated(db *gorm.DB) Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("authorization")
			split := strings.Split(authorization, "Bearer ")
			if len(split) != 2 {
				log.Println("Invalid authorization: Invalid token")
				http.Error(w, "Invalid authorization", http.StatusForbidden)
				return
			}
			tokenString := split[1]

			_, ok := jwt.ParseToken(tokenString)
			if !ok {
				log.Println("Invalid authorization: Error parsing token")
				http.Error(w, "Invalid authorization", http.StatusForbidden)
				return
			}
			hf(w, r)
		}
	}
}
