package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Не удалось найти .env файл")
	}
	key := os.Getenv("KEY")
	return &Config{
		Key: key,
	}
}
