package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config holds all the configuration settings.
	Config struct {
		App      App   `yaml:"app"`
		HTTP     HTTP  `yaml:"api"`
		Log      Log   `yaml:"logger"`
		CORS     CORS  `yaml:"cors"`
		Postgres PG    `yaml:"postgres"`
		RabbitMQ RMQ   `yaml:"rabbitmq"`
		Redis    Redis `yaml:"cache"`
	}

	// App holds application-specific settings.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP holds HTTP server settings.
	HTTP struct {
		Port int `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log holds logging settings.
	Log struct {
		Level      string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
		RollbarEnv string `yaml:"rollbar_env" env:"ROLLBAR_ENV"`
	}

	// CORS holds CORS configuration.
	CORS struct {
		AllowOrigins     []string `yaml:"allow_origins"`
		AllowMethods     []string `yaml:"allow_methods"`
		AllowHeaders     []string `yaml:"allow_headers"`
		AllowCredentials bool     `yaml:"allow_credentials"`
	}

	// PG holds PostgresSQL configuration.
	PG struct {
		Host     string `env-required:"true" yaml:"host"     env:"PG_HOST"`
		Port     string `env-required:"true" yaml:"port"     env:"PG_PORT"`
		User     string `env-required:"true" yaml:"user"     env:"PG_USER"`
		Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
		DBName   string `env-required:"true" yaml:"dbname"   env:"PG_DBNAME"`
		SSLMode  string `env-required:"true" yaml:"sslmode"  env:"PG_SSLMODE"`
		PoolMax  int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL      string
	}

	// RMQ holds RabbitMQ configuration.
	RMQ struct {
		ServerExchange string `env-required:"true" yaml:"rpc_server_exchange" env:"RMQ_RPC_SERVER"`
		ClientExchange string `env-required:"true" yaml:"rpc_client_exchange" env:"RMQ_RPC_CLIENT"`
	}

	// Redis holds Redis configuration.
	Redis struct {
		Host     string `env-required:"true" yaml:"host"     env:"REDIS_HOST"`
		Port     string `env-required:"true" yaml:"port"     env:"REDIS_PORT"`
		Password string `env-required:"true" yaml:"password" env:"REDIS_PASSWORD"`
		DB       int    `env-required:"true" yaml:"db"       env:"REDIS_DB"`
		PoolMax  int    `env-required:"true" yaml:"pool_max" env:"REDIS_POOL_MAX"`
		Address  string
	}
)

// NewConfig returns app configuration parsed from the config file and environment variables.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfg.Postgres.URL = buildURL(&cfg.Postgres)
	cfg.Redis.Address = buildAddr(&cfg.Redis)

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("environment config error: %w", err)
	}

	return cfg, nil
}

// buildURL returns a formatted URL string for PostgresSQL connection.
func buildURL(pg *PG) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s pool_max=%d",
		pg.Host, pg.Port, pg.User, pg.Password, pg.DBName, pg.SSLMode, pg.PoolMax)
}
func buildAddr(cfg *Redis) string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}
