package subscriptions_transport_http

import (
	"net/http"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_response "github.com/Vlad-6894/test_task/internal/core/transport/http/response"
	core_http_utils "github.com/Vlad-6894/test_task/internal/core/transport/http/utils"
)

type GetSubscriptionResponse SubscriptionResponseDTO

func (h *SubscribtionsHTTPHandler) GetSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	subscriptionID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "fail to get subscription id")
		return
	}

	subscription, err := h.subscriptionService.GetSubscription(ctx, subscriptionID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get subscription")
		return
	}

	startDate := core_subcription_date.ParseDateStartToString(subscription.DateStart)
	finishDate := core_subcription_date.ParseDateFinishToString(subscription.DateFinish)

	response := GetSubscriptionResponse(subscriptionDtoFromDomain(subscription, startDate, finishDate))
	responseHandler.ToJSONRsponse(response, http.StatusOK)
}
