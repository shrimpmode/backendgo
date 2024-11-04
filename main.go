package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"webserver/db"
	"webserver/middleware"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
		return
	}

	database := db.InitDB()
	db.MigrateModels(database)

	r := RegisterRoutes()

	origins := handlers.AllowedOrigins(
		[]string{os.Getenv("APP_ORIGIN")},
		// []string{"*"},
	)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	corsHandler := handlers.CORS(origins, allowedHeaders)(r)

	srv := &http.Server{
		Handler:      middleware.NewLogger(corsHandler),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
