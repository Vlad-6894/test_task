package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, w http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		w:   w,
	}
}

func (h *HTTPResponseHandler) ToJSONRsponse(
	responseBody any,
	statusCode int,
) {
	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil {
		h.log.Error("WriteHTTP response: ", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, message string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn

	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug

	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn

	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(message, zap.Error(err))

	h.errorResponse(statusCode, err, message)
}

func (h *HTTPResponseHandler) PanicResponse(p any, message string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("Panic! %v", p)

	h.log.Error(message, zap.Error(err))
	h.errorResponse(statusCode, err, message)
}

func (h *HTTPResponseHandler) errorResponse(
	statusCode int,
	err error,
	message string,
) {
	h.w.WriteHeader(statusCode)

	response := map[string]string{
		"message": message,
		"error":   err.Error(),
	}

	h.ToJSONRsponse(response, statusCode)
}
