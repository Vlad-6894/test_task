package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vlad-6894/test_task/internal/core/domain"
)

func (s *SubscriptionService) PatchSubscription(
	ctx context.Context,
	id int,
	patch domain.SubscriptionPatch,
) (domain.Subscription, error) {
	subscription, err := s.subscriptinsRepository.GetSubscription(ctx, id)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("get subscription error: %w", err)
	}

	if err := subscription.ApplyPatch(patch); err != nil {
		return domain.Subscription{}, fmt.Errorf("apply subscription patch error: %w", err)
	}

	patchedSubscription, err := s.subscriptinsRepository.PatchSubscription(ctx, id, subscription)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("subscription patch error: %w", err)
	}

	return patchedSubscription, nil
}
