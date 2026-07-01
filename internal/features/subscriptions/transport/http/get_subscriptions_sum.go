package subscriptions_transport_http

import (
	"net/http"

	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_response "github.com/Vlad-6894/test_task/internal/core/transport/http/response"
	core_http_utils "github.com/Vlad-6894/test_task/internal/core/transport/http/utils"
)

type GetSubscriptionsSumResponse int

func (h *SubscribtionsHTTPHandler) GetSubscriptionsSum(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_utils.GetUUIDPathValue(r, "user_id")
	if err != nil {
		responseHandler.ErrorResponse(err, "fail to get user_id!")
	}

	fromDate, err := core_http_utils.GetDateQueryParam(r, "from_date")
	if err != nil {
		responseHandler.ErrorResponse(err, "fail to get start_date!")
	}

	toDate, err := core_http_utils.GetDateQueryParam(r, "to_date")
	if err != nil {
		responseHandler.ErrorResponse(err, "fail to get finish_date!")
	}

	sum, err := h.subscriptionService.GetSubscriptionsSum(ctx, userID, fromDate, toDate)
	if err != nil {
		responseHandler.ErrorResponse(err, "fail to get subscriptions sum!")
	}

	response := GetSubscriptionsSumResponse(sum)

	responseHandler.ToJSONRsponse(response, http.StatusOK)
}
