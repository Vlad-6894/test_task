package subscriptions_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_response "github.com/Vlad-6894/test_task/internal/core/transport/http/response"
	core_http_utils "github.com/Vlad-6894/test_task/internal/core/transport/http/utils"
)

type GetSubscriptionsResponse []SubscriptionResponseDTO

func (h *SubscribtionsHTTPHandler) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get limit or offset query param")
		return
	}

	domains, err := h.subscriptionService.GetSubscriptions(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get subscriptions")
		return
	}

	response := GetSubscriptionsResponse(subscriptionsDTOFromdomains(domains))

	responseHandler.ToJSONRsponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get limit query param error: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get offset query param error: %w", err)
	}

	return limit, offset, nil
}
