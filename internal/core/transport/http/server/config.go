package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ConfigHTTPServer struct {
	Addr            string        `envconfig:"ADDR" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfigHTTPServer() (ConfigHTTPServer, error) {
	var config ConfigHTTPServer

	if err := envconfig.Process("HTTP", &config); err != nil {
		return ConfigHTTPServer{}, fmt.Errorf("Process envconfig HTTP: %w", err)
	}

	return config, nil
}

func NewConfigHTTPServerMust() ConfigHTTPServer {
	config, err := NewConfigHTTPServer()
	if err != nil {
		err = fmt.Errorf("Get HTTP config error: %w", err)
		panic(err)
	}

	return config
}
