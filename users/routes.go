package users

import (
	"net/http"
)

func DefineRoutes(r *http.ServeMux) {
	r.Handle("POST /users", CreateUserHandler())
}
