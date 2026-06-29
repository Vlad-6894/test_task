package subscriptions_service

import (
	"context"
	"fmt"

	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
)

func (s *SubscriptionService) DeleteSubscription(ctx context.Context, id int) error {
	if id < 0 {
		return fmt.Errorf(
			"error: negative id must not be %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if err := s.subscriptinsRepository.DeleteSubscription(ctx, id); err != nil {
		return fmt.Errorf("fail to delete subscription from repository %w", err)
	}

	return nil
}
