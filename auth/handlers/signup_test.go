package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"webserver/auth/inputs"
	"webserver/models"
)

type inputReaderMock struct {
	input inputs.SignUpInput
	err   error
}

func (i *inputReaderMock) GetInput(r *http.Request) (inputs.SignUpInput, error) {
	return i.input, i.err
}

type signUpServiceMock struct {
	verifyTokenResult  bool
	verifyTokenError   error
	hashPasswordResult string
	hashPasswordError  error
	createUserResult   error
	createUserError    error
}

func (s *signUpServiceMock) VerifyToken(token string) (bool, error) {
	return s.verifyTokenResult, s.verifyTokenError
}

func (s *signUpServiceMock) HashPassword(password string) (string, error) {
	return s.hashPasswordResult, s.hashPasswordError
}

func (s *signUpServiceMock) CreateUser(user models.User) error {
	return s.createUserError
}

func TestSignUpHandler(t *testing.T) {
	handler := &SignUpHandler{
		inputReader: &inputReaderMock{
			input: inputs.SignUpInput{
				Username:    "user1201",
				DisplayName: "user1201",
				Email:       "user1201@email.com",
				Password:    "password",
			},
		},
		signUpService: &signUpServiceMock{
			verifyTokenResult:  true,
			hashPasswordResult: "hashed_password",
			createUserResult:   nil,
		},
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/signup", nil)

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
