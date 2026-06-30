package core_postgres_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ConnectionPoolConfig struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" requred:"true"`
	Password string        `envconfig:"PASSWORD" requred:"true"`
	Database string        `envconfig:"DB" requred:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" requred:"true"`
}

func NewConnectionPoolConfig() (ConnectionPoolConfig, error) {
	var config ConnectionPoolConfig

	if err := envconfig.Process("POSTGRES", &config); err != nil {
		return ConnectionPoolConfig{}, fmt.Errorf("Process envconfig errror: %w", err)
	}

	return config, nil
}

func NewConnectionPoolConfigMust() ConnectionPoolConfig {
	config, err := NewConnectionPoolConfig()
	if err != nil {
		err = fmt.Errorf("get Postgres connection pool config error: %w", err)
		panic(err)
	}

	return config
}
