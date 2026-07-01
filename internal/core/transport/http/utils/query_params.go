package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param=%s by key=%s not a valid integer: %v:%w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return &val, nil
}

func GetDateQueryParam(r *http.Request, key string) (*core_subcription_date.YearMonth, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	date, err := core_subcription_date.ParseFinishDateFromJson(&param)
	if err != nil {
		return nil, fmt.Errorf("fail to parse date: %w", core_errors.ErrInvalidArgument)
	}

	return date, nil
}
