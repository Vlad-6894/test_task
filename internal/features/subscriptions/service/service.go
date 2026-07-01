package subscriptions_service

import (
	"context"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/Vlad-6894/test_task/internal/core/domain"
	"github.com/google/uuid"
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

	PatchSubscription(
		ctx context.Context,
		id int,
		subscription domain.Subscription,
	) (domain.Subscription, error)

	GetSubscriptionsSum(
		ctx context.Context,
		user_id uuid.UUID,
		fromDate core_subcription_date.YearMonth,
		toDate core_subcription_date.YearMonth,
	) (int, error)

	GetEarliestDate(
		ctx context.Context,
		user_id uuid.UUID,
	) (core_subcription_date.YearMonth, error)

	GetLatestSDate(
		ctx context.Context,
		user_id uuid.UUID,
	) (core_subcription_date.YearMonth, error)
}

func NewSubscriptionService(subscriptinsRepository SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		subscriptinsRepository: subscriptinsRepository,
	}
}
