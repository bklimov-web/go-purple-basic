package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")

	if key == "" {
		panic("KEY отсутствует в переменных окружения")
	}

	return &Config{
		Key: key,
	}
}
