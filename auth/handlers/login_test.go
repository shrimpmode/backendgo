package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"webserver/auth/inputs"
	"webserver/models"
)

type MockService struct {
	user     models.User
	userErr  error
	input    inputs.LoginInput
	inputErr error
	token    string
	tokenErr error
}

func (s *MockService) GetUserByEmail(email string) (models.User, error) {
	return s.user, s.userErr
}

func (s *MockService) GenerateToken(input inputs.LoginInput, user models.User) (string, error) {
	return s.token, s.tokenErr
}

type inputReader struct {
	err   error
	input inputs.LoginInput
}

func (i *inputReader) GetInput(r *http.Request) (inputs.LoginInput, error) {
	return i.input, i.err
}

func TestLoginHappyPath(t *testing.T) {
	handler := &LoginHandler{
		loginService: &MockService{
			user: models.User{
				Email:    "test@test.com",
				Password: "password",
			},
			input: inputs.LoginInput{
				Email:    "test@test.com",
				Password: "password",
			},
			token: "token",
		},
		inputReader: &inputReader{
			input: inputs.LoginInput{
				Email:    "test@test.com",
				Password: "password",
			},
			err: nil,
		},
	}

	request := httptest.NewRequest("POST", "/api/login", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, request)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestLoginInvalidUser(t *testing.T) {
	handler := &LoginHandler{
		loginService: &MockService{
			inputErr: nil,
			userErr:  errors.New("user not found"),
		},
		inputReader: &inputReader{
			err: nil,
		},
	}

	request := httptest.NewRequest("POST", "/api/login", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, request)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestLoginInvalidToken(t *testing.T) {
	handler := &LoginHandler{
		loginService: &MockService{
			tokenErr: errors.New("invalid token"),
		},
		inputReader: &inputReader{
			err: nil,
		},
	}

	request := httptest.NewRequest("POST", "/api/login", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, request)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}
}
