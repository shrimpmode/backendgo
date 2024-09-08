package handlers

type AppError struct {
	Msg string
}

var CreateChatInvalidRequest = AppError{
	Msg: "Invalid request for create chat",
}

var CreateChatInternalError = AppError{
	Msg: "Create chat internal error",
}

var CreateChatForbidden = AppError{
	Msg: "Create chat forbidden",
}
