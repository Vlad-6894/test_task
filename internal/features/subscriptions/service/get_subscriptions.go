package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vlad-6894/test_task/internal/core/domain"
	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
)

func (s *SubscriptionService) GetSubscriptions(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Subscription, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non-negative : %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non-negative : %w", core_errors.ErrInvalidArgument)
	}

	subscriptions, err := s.subscriptinsRepository.GetSubscriptions(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error get subscriptions from repository: %w", err)
	}

	return subscriptions, nil
}
