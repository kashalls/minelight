package internal

import (
	"github.com/caarlos0/env/v11"
	log "github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Dashboard   bool   `env:"DASHBOARD" envDefault:"true"`
	DryRun      bool   `env:"DRY_RUN" envDefault:"false"`
	LogFormat   string `env:"LOG_FORMAT" envDefault:"text"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	ServerHost  string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	ServerPort  int    `env:"SERVER_PORT" envDefault:"8080"`
	MetricsHost string `env:"METRICS_HOST" envDefault:"0.0.0.0"`
	MetricsPort int    `env:"METRICS_PORT" envDefault:"9090"`
}

func InitServerConfig() ServerConfig {
	cfg := ServerConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Warnf("Failed to parse environment variables: %v", err)
		return ServerConfig{}
	}

	return cfg
}
