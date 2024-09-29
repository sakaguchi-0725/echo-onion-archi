package config

import (
	"os"
)

type AppConfig struct {
	JWTSecret string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
