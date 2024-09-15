package handlers

import "net/http"

type AppError struct {
	Msg  string
	Code uint
}

const (
	ChatNameExists = "A chat with this name already exists in the server."
)

var CreateChatInvalidRequest = AppError{
	Msg: "Invalid request for create chat",
}

var CreateChatInternalError = AppError{
	Msg: "Create chat internal error",
}

var CreateChatForbidden = AppError{
	Msg: "Create chat forbidden",
}

var ErrChatNameExists = AppError{
	Msg:  ChatNameExists,
	Code: http.StatusConflict,
}
