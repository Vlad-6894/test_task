package domain

import (
	"fmt"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
	"github.com/google/uuid"
)

type Subscription struct {
	ID          int
	Version     int
	ServiceName string
	Price       int
	UserID      uuid.UUID
	DateStart   core_subcription_date.YearMonth
	DateFinish  *core_subcription_date.YearMonth
}

func NewSubscriptionCreate(
	serviceName string,
	price int,
	userID uuid.UUID,
	dateStart core_subcription_date.YearMonth,
	dateFinish *core_subcription_date.YearMonth,
) Subscription {
	return Subscription{
		ID:          UninitializedID,
		Version:     UninitializedVersion,
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		DateStart:   dateStart,
		DateFinish:  dateFinish,
	}
}

func NewSubscription(
	id int,
	version int,
	serviceName string,
	price int,
	userID uuid.UUID,
	dateStart core_subcription_date.YearMonth,
	dateFinish *core_subcription_date.YearMonth,
) Subscription {
	return Subscription{
		ID:          id,
		Version:     version,
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		DateStart:   dateStart,
		DateFinish:  dateFinish,
	}
}

func (s Subscription) Validate() error {
	dateFinishWithoutPtr := *s.DateFinish
	if s.DateStart.Year > dateFinishWithoutPtr.Year {
		return fmt.Errorf("Year date_start is later than year date_finish! %w", core_errors.ErrInvalidArgument)
	}
	if int(s.DateStart.Month) > int(dateFinishWithoutPtr.Month) {
		return fmt.Errorf("Year date_start is later than year date_finish! %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
