package main

import (
	"net/http"
	"webserver/db"
	"webserver/routes"
)

func main() {
	database := db.InitDB()

	r := routes.RegisterRoutes(database)

	http.ListenAndServe(":8080", r)
}
