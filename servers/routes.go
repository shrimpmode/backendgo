package servers

import (
	"net/http"
	"webserver/servers/handlers"
)

func DefineRoutes(r *http.ServeMux) {
	r.Handle("POST /server", handlers.Create())
}
