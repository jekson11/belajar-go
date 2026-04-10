package main

import (
	"fmt"
	"os"
)

type AppConfig struct {
	Port string
}

func LoadAppConfig() (AppConfig, error) {
	port, err := getEnv("APP_PORT")
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Port: port,
	}, nil
}

func getEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("env var %s not set", key)
	}
	return val, nil
}
