package subscriptions_transport_http

import (
	"fmt"
	"net/http"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/Vlad-6894/test_task/internal/core/domain"
	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_request "github.com/Vlad-6894/test_task/internal/core/transport/http/request"
	core_http_response "github.com/Vlad-6894/test_task/internal/core/transport/http/response"
	core_http_types "github.com/Vlad-6894/test_task/internal/core/transport/http/types"
	core_http_utils "github.com/Vlad-6894/test_task/internal/core/transport/http/utils"
)

type PatchSubscriptionRequest struct {
	Price      core_http_types.Nullable[int]    `json:"price"          swaggertype:"int" example:"400"`
	StartDate  core_http_types.Nullable[string] `json:"start_date"     swaggertype:"string" example:"07-2026"`
	FinishDate core_http_types.Nullable[string] `json:"finish_date"    swaggertype:"string" example:"08-2026"`
}

type PatchSubscriptionResponse SubscriptionResponseDTO

func (r *PatchSubscriptionRequest) Validate() error {
	if r.Price.Set {
		if r.Price.Value == nil {
			return fmt.Errorf("Price can not be NULL!")
		}
	}

	if r.StartDate.Set {
		if r.StartDate.Value == nil {
			return fmt.Errorf("Start_date can not be NULL!")
		}
	}

	return nil
}

// PatchSubscription godoc
// @Summary Изменить подписку
// @Description Изменить информацию об уже существующей подписке
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID изменяемой подписки"
// @Param request body PatchSubscriptionRequest true  "PatchSubscription тело запроса"
// @Success 200 {object} PatchSubscriptionResponse "Успешно изменённая подписка"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /subscription/{id} [patch]
func (h *SubscribtionsHTTPHandler) PatchSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	subscriptionID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "fail to get id!")
		return
	}

	var request PatchSubscriptionRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "fail to decode and validate request!")
		return
	}

	subscriptionPatch := subscriptionPatchFromRequest(request)

	subscriptionDomain, err := h.subscriptionService.PatchSubscription(ctx, subscriptionID, subscriptionPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch subscription")
		return
	}

	startDate := core_subcription_date.ParseDateStartToString(subscriptionDomain.DateStart)
	finishDate := core_subcription_date.ParseDateFinishToString(subscriptionDomain.DateFinish)

	response := PatchSubscriptionResponse(subscriptionDtoFromDomain(subscriptionDomain, startDate, finishDate))

	responseHandler.ToJSONRsponse(response, http.StatusOK)
}

func subscriptionPatchFromRequest(request PatchSubscriptionRequest) domain.SubscriptionPatch {
	return domain.SubscriptionPatch{
		Price:      request.Price.ToDomain(),
		StartDate:  request.StartDate.ToDomain(),
		FinishDate: request.FinishDate.ToDomain(),
	}
}
