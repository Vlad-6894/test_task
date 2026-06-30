package subscriptions_postgres_repository

import core_postgres_pool "github.com/Vlad-6894/test_task/internal/core/repository/postgres/pool"

type SubscriptionRepository struct {
	pool core_postgres_pool.Pool
}

func NewSubscriptionRepository(pool core_postgres_pool.Pool) *SubscriptionRepository {
	return &SubscriptionRepository{
		pool: pool,
	}
}
