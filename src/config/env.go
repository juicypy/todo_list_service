package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

func ConfigFromEnv() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	cfg.LogLevel = strings.ToLower(cfg.LogLevel)

	return cfg, err
}

type StorageConfig struct {
	Host         string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port         string `env:"DB_PORT" envDefault:"5432"`
	Username     string `env:"DB_USER" envDefault:"postgres"`
	Password     string `env:"DB_PASSWORD,required"`
	DBName       string `env:"DB_NAME" envDefault:"donut_data_db"`
	SSLMode      string `env:"DB_SSL_MODE" envDefault:"disable"`
	MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS" envDefault:"20"`
	MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS" envDefault:"20"`
}

func (c StorageConfig) DSN() string {
	dsn := "host=%s port=%s user=%s password='%s' dbname=%s sslmode=%s"
	return fmt.Sprintf(dsn,
		c.Host,
		c.Port,
		c.Username,
		c.Password,
		c.DBName,
		c.SSLMode,
	)
}

func StorageConfigFromEnv() (cfg StorageConfig, err error) {
	err = env.Parse(&cfg)
	return
}
