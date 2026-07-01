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

func (r *SubscriptionRepository) PatchSubscription(
	ctx context.Context,
	id int,
	subscription domain.Subscription,
) (domain.Subscription, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	sqlRequest := `
	UPDATE test.subscriptions
	SET 
		price=$1,
		start_date=$2,
		finish_date=$3,
		version=version+1
	WHERE id=$4 AND version=$5
	RETURNING id, version, service_name, price, user_id, start_date, finish_date;
	`

	dateStartToDB := core_subcription_date.GetDateForRepository(subscription.DateStart)
	dateFinishToDB := core_subcription_date.GetDateFinishForRepository(subscription.DateFinish)

	row := r.pool.QueryRow(
		ctxWithTime,
		sqlRequest,
		subscription.Price,
		dateStartToDB,
		dateFinishToDB,
		id,
		subscription.Version,
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
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Subscription{}, fmt.Errorf(
				"subscription with id=%d concurelly accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.Subscription{}, fmt.Errorf("scan error: %w", err)
	}

	startDate := core_subcription_date.GetStartDateFromModel(subscriptionModel.DateStart)
	finishDate := core_subcription_date.GetFinishDateFromModel(subscriptionModel.DateFinish)

	subscriptionDomain := domain.NewSubscription(
		subscriptionModel.ID,
		subscriptionModel.Version,
		subscriptionModel.ServiceName,
		subscriptionModel.Price,
		subscriptionModel.UserID,
		startDate,
		finishDate,
	)

	return subscriptionDomain, nil
}
