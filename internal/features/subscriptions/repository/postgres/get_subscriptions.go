package subscriptions_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Vlad-6894/test_task/internal/core/domain"
)

func (r *SubscriptionRepository) GetSubscriptions(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Subscription, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	sqlRequest := `
	SELECT id, version, service_name, price, user_id, start_date, finish_date FROM test.subscriptions
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.pool.Query(ctxWithTime, sqlRequest, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Select subscriptions error: %w", err)
	}
	defer rows.Close()

	var subscriptionModels []SubscriptionsModel
	for rows.Next() {
		var subscriptionModel SubscriptionsModel

		if err := rows.Scan(
			&subscriptionModel.ID,
			&subscriptionModel.Version,
			&subscriptionModel.ServiceName,
			&subscriptionModel.Price,
			&subscriptionModel.UserID,
			&subscriptionModel.DateStart,
			&subscriptionModel.DateFinish,
		); err != nil {
			return nil, fmt.Errorf("Scan subscription error: %w", err)
		}

		subscriptionModels = append(subscriptionModels, subscriptionModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows error: %w", err)
	}

	subscriptionDomains := subscriptionDomansFromModels(subscriptionModels)

	return subscriptionDomains, nil
}
