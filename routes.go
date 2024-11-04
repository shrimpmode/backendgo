package main

import (
	"net/http"
	"webserver/auth"
	"webserver/users"
)

func RegisterRoutes() *http.ServeMux {

	r := http.NewServeMux()
	auth.DefineRoutes(r)

	apiRouter := http.NewServeMux()
	users.DefineRoutes(apiRouter)

	r.Handle("/api/", http.StripPrefix("/api", apiRouter))

	return r
}
