package subscriptions_postgres_repository

import (
	"time"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/google/uuid"
)

type SubscriptionsCreateRequestModel struct {
	ServiceName string
	Price       int
	UserID      uuid.UUID
	DateStart   time.Time
	DateFinish  *time.Time
}

type SubscriptionsModel struct {
	ID          int
	Version     int
	ServiceName string
	Price       int
	UserID      uuid.UUID
	DateStart   time.Time
	DateFinish  *time.Time
}

func NewSubscriptionsCreateRequestModel(
	serviceName string,
	price int,
	userID uuid.UUID,
	dateStart core_subcription_date.YearMonth,
	dateFinish *core_subcription_date.YearMonth,
) SubscriptionsCreateRequestModel {
	finishDate := time.Date(dateFinish.Year, dateFinish.Month, 1, 0, 0, 0, 0, time.UTC)
	return SubscriptionsCreateRequestModel{
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		DateStart:   time.Date(dateStart.Year, dateStart.Month, 1, 0, 0, 0, 0, time.UTC),
		DateFinish:  &finishDate,
	}
}
