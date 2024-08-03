package store

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func NewStore() *sessions.CookieStore {
	godotenv.Load()
	key := []byte(os.Getenv("SESSION_KEY"))
	store := sessions.NewCookieStore(key)
	return store
}
