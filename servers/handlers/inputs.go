package handlers

type CreateServerRequest struct {
	Name string `json:"name" validate:"required"`
}
