package inputs

type SignUpInput struct {
	Email           string `validate:"required"`
	Password        string `validate:"required"`
	PasswordConfirm string `validate:"required"`
	UserName        string `validate:"required"`
	DisplayName     string `validate:"required"`
	Token           string
}

type RecaptchaResponse struct {
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes,omitempty"`
	Success     bool     `json:"success"`
}
