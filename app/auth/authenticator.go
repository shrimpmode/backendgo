package auth

import (
	"net/http"
	"webserver/models"
)

type Authenticator interface {
	GetAuthenticatedUser(r *http.Request) (*models.User, bool)
}
