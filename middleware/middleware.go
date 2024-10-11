package middleware

import (
	"net/http"
)

type MiddlewareFunc func(http.Handler) http.Handler

func ApplyMiddlewares(h http.Handler, middlewares ...MiddlewareFunc) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
