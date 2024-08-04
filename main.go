package main

import (
	"net/http"
	"webserver/db"
	"webserver/routes"
	"webserver/store"
)

func main() {
	store := store.NewStore()

	database := db.InitDB()
	db.MigrateModles(database)

	r := routes.RegisterRoutes(database, store)

	http.ListenAndServe(":8080", r)
}
