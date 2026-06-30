package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dto any) error {
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		return fmt.Errorf("Decode DTO error: %v :  %w", err, core_errors.ErrInvalidArgument)
	}

	v, ok := dto.(validatable)
	if ok {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("request validate error : %v : %w", err, core_errors.ErrInvalidArgument)
		}
	} else {
		if err := requestValidator.Struct(dto); err != nil {
			return fmt.Errorf("request validate error : %v : %w", err, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}
