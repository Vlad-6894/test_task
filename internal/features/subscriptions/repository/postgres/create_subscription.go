package subscriptions_postgres_repository

import (
	"context"
	"fmt"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/Vlad-6894/test_task/internal/core/domain"
)

func (r *SubscriptionRepository) CreateSubscription(
	ctx context.Context,
	subscription domain.Subscription,
) (domain.Subscription, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	sqlRequest := `
	INSERT INTO test.subscriptions (service_name, price, user_id, start_date, finish_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, version, service_name, price, user_id, start_date, finish_date;
	`
	requestModel := NewSubscriptionsCreateRequestModel(
		subscription.ServiceName,
		subscription.Price,
		subscription.UserID,
		subscription.DateStart,
		subscription.DateFinish,
	)

	row := r.pool.QueryRow(
		ctxWithTime,
		sqlRequest,
		requestModel.ServiceName,
		requestModel.Price,
		requestModel.UserID,
		requestModel.DateStart,
		requestModel.DateFinish,
	)

	var subscriptionModel SubscriptionsModel
	if err := row.Scan(
		&subscriptionModel.ID,
		&subscriptionModel.Version,
		&subscriptionModel.ServiceName,
		&subscriptionModel.Price,
		&subscriptionModel.UserID,
		&subscriptionModel.DateStart,
		&subscriptionModel.DateFinish,
	); err != nil {
		return domain.Subscription{}, fmt.Errorf("scan error: %w", err)
	}

	startDate := core_subcription_date.GetStartDateFromModel(subscriptionModel.DateStart)
	finishDate := core_subcription_date.GetFinishDateFromModel(subscriptionModel.DateFinish)
	subscriptionDomain := domain.NewSubscription(
		subscriptionModel.ID,
		subscription.Version,
		subscriptionModel.ServiceName,
		subscriptionModel.Price,
		subscriptionModel.UserID,
		startDate,
		finishDate,
	)

	return subscriptionDomain, nil
}
