package users

import (
	"encoding/json"
	"net/http"
	"webserver/app"
	"webserver/app/routehandler"
	"webserver/auth/passwords"
	"webserver/db"
	"webserver/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request, ctx *app.Context) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var createUserInput CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&createUserInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(&createUserInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password, _ := passwords.HashPassword(createUserInput.Password)

	user := models.User{
		Email:    createUserInput.Email,
		Username: createUserInput.Username,
		Password: password,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func CreateUserHandler() http.Handler {
	h := &Handler{DB: db.GetDB()}

	return routehandler.NewHandler(h)
}
