package users

import (
	"net/http"
	"webserver/users/handlers"
)

func DefineRoutes(r *http.ServeMux) {
	r.Handle("POST /users", handlers.Create())
}
