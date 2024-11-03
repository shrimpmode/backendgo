package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/auth/inputs"

	"github.com/go-playground/validator/v10"
)

type LoginHandler struct {
	loginService LoginService
	inputReader  InputReader[inputs.LoginInput]
}

type InputReader[T any] interface {
	GetInput(r *http.Request) (T, error)
}

type Input[T any] struct {
	input T
}

func (i *Input[T]) GetInput(r *http.Request) (T, error) {
	if err := json.NewDecoder(r.Body).Decode(&i.input); err != nil {
		return i.input, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(i.input); err != nil {
		return i.input, err
	}

	return i.input, nil
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	input, err := h.inputReader.GetInput(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.loginService.GetUserByEmail(input.Email)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.loginService.GenerateToken(input, user)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	data := inputs.TokenResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{
		loginService: NewLoginService(),
		inputReader:  &Input[inputs.LoginInput]{},
	}
}
