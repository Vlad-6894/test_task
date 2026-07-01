package subscriptions_postgres_repository

import (
	"context"
	"fmt"
	"time"

	core_subcription_date "github.com/Vlad-6894/test_task/internal/core/date/subsription"
	"github.com/google/uuid"
)

func (r *SubscriptionRepository) GetSubscriptionsSum(
	ctx context.Context,
	user_id uuid.UUID,
	fromDate core_subcription_date.YearMonth,
	toDate core_subcription_date.YearMonth,
) (int, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	newFromDate := core_subcription_date.GetDateForRepository(fromDate)
	newToDate := core_subcription_date.GetDateForRepository(toDate)

	sqlRequest := `
	SELECT SUM(price) FROM test.subscriptions
	WHERE start_date >= $1 AND start_date <= $2
	GROUP BY user_id
	HAVING user_id = $3;
	`

	row := r.pool.QueryRow(ctxWithTime, sqlRequest, newFromDate, newToDate, user_id)
	var sum int

	if err := row.Scan(&sum); err != nil {
		return 0, fmt.Errorf("fail to scan: %w", err)
	}

	return sum, nil
}

func (r *SubscriptionRepository) GetEarliestDate(
	ctx context.Context,
	user_id uuid.UUID,
) (core_subcription_date.YearMonth, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	sqlRequest := `
	SELECT MIN(start_date) FROM test.subscriptions
	WHERE user_id = $1;
	`

	row := r.pool.QueryRow(ctxWithTime, sqlRequest, user_id)
	var date time.Time

	if err := row.Scan(&date); err != nil {
		return core_subcription_date.YearMonth{}, fmt.Errorf("fail to scan: %w", err)
	}

	newDate := core_subcription_date.GetStartDateFromModel(date)

	return newDate, nil
}

func (r *SubscriptionRepository) GetLatestSDate(
	ctx context.Context,
	user_id uuid.UUID,
) (core_subcription_date.YearMonth, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	sqlRequest := `
	SELECT MAX(start_date) FROM test.subscriptions
	WHERE user_id = $1;
	`

	row := r.pool.QueryRow(ctxWithTime, sqlRequest, user_id)
	var date time.Time

	if err := row.Scan(&date); err != nil {
		return core_subcription_date.YearMonth{}, fmt.Errorf("fail to scan: %w", err)
	}

	newDate := core_subcription_date.GetStartDateFromModel(date)

	return newDate, nil
}
