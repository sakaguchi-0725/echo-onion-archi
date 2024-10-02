package db

import (
	"log"

	"github.com/sakaguchi-0725/echo-onion-arch/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDB(config *config.TestDBConfig) (*gorm.DB, error) {
	dsn := config.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the test database!")
	return db, nil
}
