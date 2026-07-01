package subscriptions_transport_http

import (
	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/Vlad-6894/test_task/internal/core/domain"
	"github.com/google/uuid"
)

type SubscriptionResponseDTO struct {
	ID          int       `json:"id"              example:"1"`
	Version     int       `json:"version"         example:"2"`
	ServiceName string    `json:"service_name"    example:"Yandex Plus"`
	Price       int       `json:"price"           example:"400"`
	UserID      uuid.UUID `json:"user_id"         example:"60601ef-ddrsa-ggg34464"`
	StartDate   string    `json:"start_date"      example:"07-2026"`
	FinishDate  *string   `json:"finish_date"     example:"08-2026"`
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
