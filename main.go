package main

import (
	"log"
	"net/http"
	"os"
	"webserver/db"
	"webserver/routes"
	"webserver/store"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
		return
	}

	store := store.NewStore()

	database := db.InitDB()
	db.MigrateModles(database)

	r := routes.RegisterRoutes(database, store)

	origins := handlers.AllowedOrigins(
		[]string{os.Getenv("APP_ORIGIN")},
	)

	http.ListenAndServe(":8080", handlers.CORS(
		origins,
	)(r))
}
