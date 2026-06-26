package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (h *HTTPResponseHandler) PanicResponse(p any, message string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("Panic! %v", p)

	h.log.Error(message, zap.Error(err))
	h.w.WriteHeader(statusCode)

	response := map[string]string{
		"message": message,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.w).Encode(response); err != nil {
		h.log.Error("Write HTTP response error", zap.Error(err))
	}
}
