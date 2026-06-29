package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vlad-6894/test_task/internal/core/domain"
)

func (s *SubscriptionService) CreateSubscription(
	ctx context.Context,
	subscription domain.Subscription,
) (domain.Subscription, error) {
	if err := subscription.Validate(); err != nil {
		return domain.Subscription{}, fmt.Errorf("Error validate subscription: %w", err)
	}

	subscription, err := s.subscriptinsRepository.CreateSubscription(ctx, subscription)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("Error create subscription: %w", err)
	}

	return subscription, nil
}
