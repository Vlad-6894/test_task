package subscriptions_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
)

func (r *SubscriptionRepository) DeleteSubscription(
	ctx context.Context,
	id int,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	sqlRequest := `
	DELETE FROM test.subscriptions
	WHERE id = $1;
	`

	cmndTag, err := r.pool.Exec(ctxWithTime, sqlRequest, id)
	if err != nil {
		return fmt.Errorf("exec query error: %w", err)
	}

	if cmndTag.RowsAffected() == 0 {
		return fmt.Errorf("subscription with id=%d: %w", id, core_errors.ErrNotFound)
	}

	return nil
}
