package handlers

type AppError struct {
	Message string
	Code    uint
}

func (e AppError) Error() string {
	return e.Message
}

// var CreateChatInvalidRequest = AppError{
// 	Msg: "Invalid request for create chat",
// }
