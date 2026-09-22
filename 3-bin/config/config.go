package config

import (
	"errors"
	"os"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	key := os.Getenv("APP_KEY")
	if key == "" {
		return nil, errors.New("No key provided!")
	}
	return &Config{
		Key: key,
	}, nil
}
