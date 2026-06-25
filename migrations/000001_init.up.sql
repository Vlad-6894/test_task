CREATE SCHEMA test;

CREATE TABLE test.subscriptions (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    service_name VARCHAR(100) NOT NULL,
    price BIGINT NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    finish_date DATE
);