package main

import (
	"net/http"
	"webserver/db"
	"webserver/routes"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	database := db.InitDB()

	r := routes.RegisterRoutes(database, store)

	http.ListenAndServe(":8080", r)
}
