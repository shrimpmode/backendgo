package inputs

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
