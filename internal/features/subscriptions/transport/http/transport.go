package subscriptions_transport_http

import (
	"context"
	"net/http"

	"github.com/Vlad-6894/test_task/internal/core/domain"
	core_http_server "github.com/Vlad-6894/test_task/internal/core/transport/http/server"
)

type SubscribtionsHTTPHandler struct {
	subscriptionService SubscriptionsService
}

type SubscriptionsService interface {
	CreateSubscription(ctx context.Context, subscription domain.Subscription) (domain.Subscription, error)
}

func NewSubscriptionsHTTPHandler(subscriptionService SubscriptionsService) *SubscribtionsHTTPHandler {
	return &SubscribtionsHTTPHandler{
		subscriptionService: subscriptionService,
	}
}

func (h *SubscribtionsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/subscription",
			Handler: h.Create,
		},
	}
}
