package inputs

type SignUpInput struct {
	Email           string `validate:"required"`
	Password        string `validate:"required"`
	PasswordConfirm string `validate:"required"`
}
