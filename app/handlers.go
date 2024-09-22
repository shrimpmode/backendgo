package app

import (
	"net/http"
	"webserver/models"
)

type AppHandler interface {
	SetUser(*models.User)
	ServeHTTP(http.ResponseWriter, *http.Request)
}
