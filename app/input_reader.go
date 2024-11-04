package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

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
