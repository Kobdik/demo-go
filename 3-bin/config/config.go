package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Ошибка загрузки .env файла")
	}
	key := os.Getenv("APP_KEY")
	if key == "" {
		return nil, errors.New("No key provided!")
	}
	return &Config{
		Key: key,
	}, nil
}
