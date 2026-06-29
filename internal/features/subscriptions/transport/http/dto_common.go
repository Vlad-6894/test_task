package subscriptions_transport_http

import (
	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/Vlad-6894/test_task/internal/core/domain"
	"github.com/google/uuid"
)

type SubscriptionResponseDTO struct {
	ID          int       `json:"id"`
	Version     int       `json:"version"`
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   string    `json:"start_date"`
	FinishDate  *string   `json:"finish_date"`
}

func subscriptionDtoFromDomain(subscription domain.Subscription, startDate string, finishDate *string) SubscriptionResponseDTO {

	return SubscriptionResponseDTO{
		ID:          subscription.ID,
		Version:     subscription.Version,
		ServiceName: subscription.ServiceName,
		Price:       subscription.Price,
		UserID:      subscription.UserID,
		StartDate:   startDate,
		FinishDate:  finishDate,
	}
}

func subscriptionsDTOFromdomains(subscriptions []domain.Subscription) []SubscriptionResponseDTO {
	subscriptionsDTO := make([]SubscriptionResponseDTO, len(subscriptions))

	for i, subscription := range subscriptions {
		startDate := core_subcription_date.ParseDateStartToString(subscription.DateStart)
		finishDate := core_subcription_date.ParseDateFinishToString(subscription.DateFinish)

		dto := subscriptionDtoFromDomain(subscription, startDate, finishDate)
		subscriptionsDTO[i] = dto
	}

	return subscriptionsDTO
}
