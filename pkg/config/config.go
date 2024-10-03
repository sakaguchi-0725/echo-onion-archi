package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	App *AppConfig
	DB  *DBConfig
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}
}

func NewConfig() *Config {
	return &Config{
		App: NewAppConfig(),
		DB:  NewDBConfig(),
	}
}
