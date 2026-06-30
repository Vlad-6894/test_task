package subscriptions_transport_http

import (
	"net/http"

	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_response "github.com/Vlad-6894/test_task/internal/core/transport/http/response"
	"github.com/google/uuid"
)

type PatchSubscriptionRequest struct {
	ServiceName string    `json:"service_name" validate:"required"`
	Price       int       `json:"price" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
}

func (h *SubscribtionsHTTPHandler) PatchSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
}
