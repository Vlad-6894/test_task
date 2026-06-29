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
	ServiceName string                           `json:"service_name" validate:"required"`
	Price       int                              `json:"price" validate:"required"`
	UserID      uuid.UUID                        `json:"user_id" validate:"required"`
	StartDate   core_subcription_date.YearMonth  `json:"start_date" validate:"required"`
	FinishDate  *core_subcription_date.YearMonth `json:"finish_date"`
}

type CreateSubscriptionRequestDTOWithoutDate struct {
	ServiceName string    `json:"service_name" validate:"required"`
	Price       int       `json:"price" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
}

type CreateSubscriptionResponseDTO SubscriptionResponseDTO

func NewCreateSubscriptionRequestDTO(
	serviceName string,
	price int,
	userID uuid.UUID,
	startDate core_subcription_date.YearMonth,
	finishDate *core_subcription_date.YearMonth,
) CreateSubscriptionRequestDTO {
	return CreateSubscriptionRequestDTO{
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		StartDate:   startDate,
		FinishDate:  finishDate,
	}
}

func (h *SubscribtionsHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Info("Create Subscription handler start")

	var request CreateSubscriptionRequestDTOWithoutDate
	var requestMap map[string]any

	if err := core_http_request.DecodeAndValidateRequest(r, &requestMap); err != nil {
		responseHandler.ErrorResponse(err, "failed decode http request!")
		return
	}

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed decode http request!")
		return
	}

	startDate, err := core_subcription_date.ParseStartDateFromJson(requestMap)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get start date")
		return
	}

	finishDate, err := core_subcription_date.ParseFinishDateFromJson(requestMap)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get start date")
		return
	}

	requestDto := NewCreateSubscriptionRequestDTO(
		request.ServiceName,
		request.Price,
		request.UserID,
		startDate,
		finishDate,
	)

	subscriptionDomain := domainFromDTO(requestDto)
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

func domainFromDTO(dto CreateSubscriptionRequestDTO) domain.Subscription {
	return domain.NewSubscriptionCreate(dto.ServiceName, dto.Price, dto.UserID, dto.StartDate, dto.FinishDate)
}
