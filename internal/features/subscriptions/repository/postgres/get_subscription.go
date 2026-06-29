package subscriptions_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/Vlad-6894/test_task/internal/core/domain"
	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *SubscriptionRepository) GetSubscription(
	ctx context.Context,
	id int,
) (domain.Subscription, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	sqlRequest := `
	SELECT id, version, service_name, price, user_id, start_date, finish_date FROM test.subscriptions
	WHERE id = $1;
	`

	row := r.pool.QueryRow(ctxWithTime, sqlRequest, id)

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
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Subscription{}, fmt.Errorf("subscription with id=%d: %w", id, core_errors.ErrNotFound)
		}

		return domain.Subscription{}, fmt.Errorf("scan error %w", err)
	}

	startDate := core_subcription_date.GetStartDateFromModel(subscriptionModel.DateStart)
	finishDate := core_subcription_date.GetFinishDateFromModel(subscriptionModel.DateFinish)

	domain := domain.NewSubscription(
		subscriptionModel.ID,
		subscriptionModel.Version,
		subscriptionModel.ServiceName,
		subscriptionModel.Price,
		subscriptionModel.UserID,
		startDate,
		finishDate,
	)

	return domain, nil
}
