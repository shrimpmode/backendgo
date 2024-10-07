package app

import (
	"net/http"
	"webserver/models"
)

type AppHandler interface {
	SetUser(*models.User)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Context struct {
	User *models.User
}

func (ctx *Context) SetUser(user *models.User) {
	ctx.User = user
}

func NewContext() *Context {
	return &Context{}
}

type Handler interface {
	Handle(http.ResponseWriter, *http.Request, *Context)
}
