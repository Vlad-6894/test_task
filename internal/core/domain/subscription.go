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

type SubscriptionPatch struct {
	Price      Nullable[int]
	StartDate  Nullable[string]
	FinishDate Nullable[string]
}

func (p *SubscriptionPatch) Validate() error {
	if p.Price.Set && p.Price.Value == nil {
		return fmt.Errorf("Price can not be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	if p.StartDate.Set && p.StartDate.Value == nil {
		return fmt.Errorf("Start_date can not be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (s *Subscription) ApplyPatch(patch SubscriptionPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate patch error: %w", err)
	}

	tmp := *s

	if patch.Price.Set {
		tmp.Price = *patch.Price.Value
	}

	if patch.StartDate.Set {
		startDate, err := core_subcription_date.ParseStartDateFromJson(*patch.StartDate.Value)
		if err != nil {
			return fmt.Errorf("parse startDate error: %w", err)
		}

		tmp.DateStart = startDate
	}

	if patch.FinishDate.Set {
		finishDate, err := core_subcription_date.ParseFinishDateFromJson(patch.FinishDate.Value)
		if err != nil {
			return fmt.Errorf("parse finishDate error: %w", err)
		}

		tmp.DateFinish = finishDate
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched subscription error: %w", err)
	}

	*s = tmp

	return nil
}
