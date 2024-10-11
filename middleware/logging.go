package middleware

import (
	"log"
	"net/http"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		f(w, r)
	}
}

type Logger struct {
	handler http.Handler
}

func (m *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	m.handler.ServeHTTP(w, r)
}

func NewLogger(handler http.Handler) http.Handler {
	return &Logger{handler: handler}
}
