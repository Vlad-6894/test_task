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
	GetSubscriptions(ctx context.Context, limit *int, offset *int) ([]domain.Subscription, error)
	GetSubscription(ctx context.Context, id int) (domain.Subscription, error)
	DeleteSubscription(ctx context.Context, id int) error
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
		{
			Method:  http.MethodGet,
			Path:    "/subscription",
			Handler: h.GetSubscriptions,
		},
		{
			Method:  http.MethodGet,
			Path:    "/subscription/{id}",
			Handler: h.GetSubscription,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/subscriptions/{id}",
			Handler: h.DeleteSubscription,
		},
	}
}
