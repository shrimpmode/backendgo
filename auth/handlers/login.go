package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/app"
	"webserver/auth/inputs"
)

type LoginHandler struct {
	loginService LoginService
	inputReader  app.InputReader[inputs.LoginInput]
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
		inputReader:  &app.Input[inputs.LoginInput]{},
	}
}
