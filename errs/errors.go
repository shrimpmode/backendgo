package errs

import "net/http"

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	AuthUserError = AppError{
		Message: "User is not authenticated",
		Code:    http.StatusUnauthorized,
	}
	AppErrBadRequest = AppError{
		"Invalid input",
		http.StatusBadRequest,
	}
)

func NewAppErr(msg string, code int) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
	}
}
