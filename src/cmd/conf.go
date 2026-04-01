package main

import (
	"fmt"
	"os"

	"belajar-go/src/config/database"
	"belajar-go/src/config/logger"
	"belajar-go/src/config/query"
	"belajar-go/src/config/server"
	"belajar-go/src/config/tracer"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Server   server.ServerOptions     `yaml:"server"`
	Logger   logger.LoggerOptions     `yaml:"logger"`
	Postgres database.DatabaseOptions `yaml:"postgres"`
	MySQL    database.DatabaseOptions `yaml:"mysql"`
	Queries  query.QueriesOptions     `yaml:"queries"`
	Gin      server.GinOptions        `yaml:"gin"`
	Tracer   tracer.TracerOptions     `yaml:"tracer"`
}

func InitConfig() (*Config, error) {
	configPath := "config.yaml"

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Override with environment variables if present
	overrideWithEnv(&cfg)

	return &cfg, nil
}

func overrideWithEnv(cfg *Config) {
	if val := os.Getenv("SERVER_PORT"); val != "" {
		cfg.Server.Port = parseInt(val, cfg.Server.Port)
	}

	if val := os.Getenv("LOG_LEVEL"); val != "" {
		cfg.Logger.Level = val
	}

	if val := os.Getenv("DB_HOST"); val != "" {
		cfg.Postgres.Host = val
	}

	if val := os.Getenv("DB_PORT"); val != "" {
		cfg.Postgres.Port = parseInt(val, cfg.Postgres.Port)
	}

	if val := os.Getenv("DB_USER"); val != "" {
		cfg.Postgres.User = val
	}

	if val := os.Getenv("DB_PASSWORD"); val != "" {
		cfg.Postgres.Password = val
	}

	if val := os.Getenv("DB_NAME"); val != "" {
		cfg.Postgres.DBName = val
	}

	if val := os.Getenv("TRACER_ENDPOINT"); val != "" {
		cfg.Tracer.Endpoint = val
	}
}

func parseInt(s string, defaultVal int) int {
	var val int
	if _, err := fmt.Sscanf(s, "%d", &val); err == nil {
		return val
	}

	return defaultVal
}
