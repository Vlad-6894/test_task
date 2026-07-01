package subscriptions_transport_http

import (
	"net/http"

	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_response "github.com/Vlad-6894/test_task/internal/core/transport/http/response"
	core_http_utils "github.com/Vlad-6894/test_task/internal/core/transport/http/utils"
)

// DeleteSubscription godoc
// @Summary Удалить подписку
// @Description Удалить подписку по её id
// @Tags subscriptions
// @Param id path int true "ID подписки, которая удаляется"
// @Success 204  "Успешное удаление подписки"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /subscription/{id} [delete]
func (h *SubscribtionsHTTPHandler) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	subscriptionID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "fail to get subscription id")
		return
	}

	if err := h.subscriptionService.DeleteSubscription(ctx, subscriptionID); err != nil {
		responseHandler.ErrorResponse(err, "fail to delete subscription")
		return
	}

	responseHandler.NoContentResponse()
}
