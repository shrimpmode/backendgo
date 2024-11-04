package app

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

type User struct {
	Name string `json:"name" validate:"required;string"`
}

func TestInputReader(t *testing.T) {

	data := &User{
		Name: "John Doe",
	}

	jsonData, _ := json.Marshal(data)

	request := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonData))

	input := Input[*User]{input: data}

	_, err := input.GetInput(request)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestInputReaderBodyError(t *testing.T) {
	request := httptest.NewRequest("POST", "/", nil)

	input := Input[*User]{input: &User{}}

	_, err := input.GetInput(request)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestInputReaderValidationError(t *testing.T) {
	data := &User{}
	jsonData, _ := json.Marshal(data)
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonData))

	input := Input[*User]{input: data}

	_, err := input.GetInput(r)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
