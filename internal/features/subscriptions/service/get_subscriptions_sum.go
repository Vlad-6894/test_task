package subscriptions_service

import (
	"context"
	"fmt"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/google/uuid"
)

func (s *SubscriptionService) GetSubscriptionsSum(
	ctx context.Context,
	user_id *uuid.UUID,
	fromDate *core_subcription_date.YearMonth,
	toDate *core_subcription_date.YearMonth,
) (int, error) {
	var fromDateNewTo core_subcription_date.YearMonth
	var ToDateNewTo core_subcription_date.YearMonth

	if fromDate == nil {
		fromDateNew, err := s.subscriptinsRepository.GetEarliestDate(ctx, *user_id)
		if err != nil {
			return 0, fmt.Errorf("fail to get earliest date! %w", err)
		}
		fromDateNewTo = fromDateNew
	}

	if toDate == nil {
		toDateNew, err := s.subscriptinsRepository.GetLatestSDate(ctx, *user_id)
		if err != nil {
			return 0, fmt.Errorf("fail to get latest date! %w", err)
		}
		ToDateNewTo = toDateNew
	}

	sum, err := s.subscriptinsRepository.GetSubscriptionsSum(ctx, *user_id, fromDateNewTo, ToDateNewTo)
	if err != nil {
		return 0, fmt.Errorf("fail to get sum from repository %w", err)
	}

	return sum, nil
}
