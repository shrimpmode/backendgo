package main

import (
	"log"
	"net/http"
	"time"
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
		// []string{os.Getenv("APP_ORIGIN")},
		[]string{"*"},
	)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	srv := &http.Server{
		Handler:      handlers.CORS(origins, allowedHeaders)(r),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
