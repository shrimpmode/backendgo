package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/auth/inputs"
)

type LoginHandler struct {
	loginService LoginServiceInterface
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	loginInput, err := h.loginService.GetInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.loginService.GetUser(loginInput.Email)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.loginService.GenerateToken(loginInput, user)
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

func Login() http.Handler {
	return &LoginHandler{
		loginService: &LoginService{},
	}
}
