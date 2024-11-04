package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"webserver/app"
	"webserver/models"
)

type serviceMock struct {
	server    *models.Server
	createErr error
}

func (s *serviceMock) Create(server *models.Server) error {
	return s.createErr
}

func NewServiceMock(server *models.Server, createErr error) ServerService {
	return &serviceMock{server, createErr}
}

func TestCreateServerHandler(t *testing.T) {
	input := &CreateServerRequest{Name: "Test"}
	jsonData, _ := json.Marshal(input)

	r := httptest.NewRequest(http.MethodPost, "/server", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()

	server := &models.Server{
		Name:    "Test",
		OwnerID: 1,
	}

	createServerHandler := &ServerHandler{NewServiceMock(server, nil)}

	createServerHandler.Handle(
		w,
		r,
		&app.Context{User: &models.User{}},
	)

	want := http.StatusCreated

	if w.Code != want {
		t.Errorf("Expected code: %d. Got: %d", want, w.Code)
	}
}
