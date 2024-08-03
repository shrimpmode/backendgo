package middleware

import (
	"net/http"

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
