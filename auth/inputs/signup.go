package inputs

type SignUpInput struct {
	Email           string `validate:"required"`
	Password        string `validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
	Username        string `json:"username" validate:"required"`
	DisplayName     string `json:"display_name" validate:"required"`
	Token           string
}

type RecaptchaResponse struct {
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes,omitempty"`
	Success     bool     `json:"success"`
}
