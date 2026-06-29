package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vlad-6894/test_task/internal/core/domain"
	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
)

func (s *SubscriptionService) GetSubscription(ctx context.Context, id int) (domain.Subscription, error) {
	if id < 0 {
		return domain.Subscription{}, fmt.Errorf(
			"error: negative id must not be %w",
			core_errors.ErrInvalidArgument,
		)
	}

	subscription, err := s.subscriptinsRepository.GetSubscription(ctx, id)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("get subscription from repository error: %w", err)
	}

	return subscription, nil
}
