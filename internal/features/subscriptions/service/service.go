package subscriptions_service

import (
	"context"

	"github.com/Vlad-6894/test_task/internal/core/domain"
)

type SubscriptionService struct {
	subscriptinsRepository SubscriptionRepository
}

type SubscriptionRepository interface {
	CreateSubscription(
		ctx context.Context,
		subscription domain.Subscription,
	) (domain.Subscription, error)

	GetSubscriptions(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Subscription, error)

	GetSubscription(
		ctx context.Context,
		id int,
	) (domain.Subscription, error)

	DeleteSubscription(
		ctx context.Context,
		id int,
	) error
}

func NewSubscriptionService(subscriptinsRepository SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		subscriptinsRepository: subscriptinsRepository,
	}
}
