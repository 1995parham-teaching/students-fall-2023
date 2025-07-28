package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type StudentCreate struct {
	Name   string `json:"name,omitempty"   validate:"required,alphaunicode"`
	Family string `json:"family,omitempty" validate:"required,alphaunicode"`
}

func (sc StudentCreate) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(sc)
	if err != nil {
		return fmt.Errorf("create request validation failed %w", err)
	}

	return nil
}
