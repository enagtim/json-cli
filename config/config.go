package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	apiKey := os.Getenv("KEY")
	if apiKey == "" {
		log.Fatal("Переменная окружения KEY не установлена")
	}
	return &Config{
		Key: apiKey,
	}
}
