package subscriptions_transport_http

import (
	"net/http"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/Vlad-6894/test_task/internal/core/domain"
	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_request "github.com/Vlad-6894/test_task/internal/core/transport/http/request"
	core_http_response "github.com/Vlad-6894/test_task/internal/core/transport/http/response"
	"github.com/google/uuid"
)

type CreateSubscriptionRequestDTO struct {
	ServiceName string    `json:"service_name" validate:"required"      example:"Yandex Plus"`
	Price       int       `json:"price" validate:"required"             example:"400"`
	UserID      uuid.UUID `json:"user_id" validate:"required"           example:"60601ef-ddrsa-ggg34464"`
	StartDate   string    `json:"start_date" validate:"required"        example:"07-2026"`
	FinishDate  *string   `json:"finish_date"                           example:"08-2026"`
}

type CreateSubscriptionResponseDTO SubscriptionResponseDTO

// Create godoc
// @Summary Создать подписку
// @Description Создать новую подписку в системе
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body CreateSubscriptionRequestDTO true "CreateSubscriptionRequestDTO тело запроса"
// @Success 201 {object} CreateSubscriptionResponseDTO "Успешно созданная подписка"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /subscription [post]
func (h *SubscribtionsHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Info("Create Subscription handler start")

	var request CreateSubscriptionRequestDTO

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed decode http request!")
		return
	}

	startDate, err := core_subcription_date.ParseStartDateFromJson(request.StartDate)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get start date")
		return
	}

	finishDate, err := core_subcription_date.ParseFinishDateFromJson(request.FinishDate)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get start date")
		return
	}

	subscriptionDomain := domainFromDTO(request, startDate, finishDate)
	subscriptionDomain, err = h.subscriptionService.CreateSubscription(ctx, subscriptionDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create")
		return
	}

	dateStart := core_subcription_date.ParseDateStartToString(subscriptionDomain.DateStart)
	dateFinish := core_subcription_date.ParseDateFinishToString(subscriptionDomain.DateFinish)

	response := CreateSubscriptionResponseDTO(subscriptionDtoFromDomain(subscriptionDomain, dateStart, dateFinish))

	responseHandler.ToJSONRsponse(response, http.StatusCreated)
}

func domainFromDTO(
	dto CreateSubscriptionRequestDTO,
	startDate core_subcription_date.YearMonth,
	finishDate *core_subcription_date.YearMonth,
) domain.Subscription {
	return domain.NewSubscriptionCreate(dto.ServiceName, dto.Price, dto.UserID, startDate, finishDate)
}
