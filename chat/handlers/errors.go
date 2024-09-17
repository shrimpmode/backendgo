package handlers

import "net/http"

type AppError struct {
	Message string
	Code    int
}

func (e AppError) Error() string {
	return e.Message
}

var (
	AppErrBadRequest = AppError{
		"Invalid input",
		http.StatusBadRequest,
	}
)
