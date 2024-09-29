package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type TestDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func NewTestDBConfig() *TestDBConfig {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(".env is not found")
	}

	return &TestDBConfig{
		Host:     "5432",
		Port:     os.Getenv("TEST_DB_PORT"),
		User:     os.Getenv("TEST_DB_USER"),
		Password: os.Getenv("TEST_DB_PASSWORD"),
		Name:     os.Getenv("TEST_DB_NAME"),
	}
}

func (config *TestDBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)
}
