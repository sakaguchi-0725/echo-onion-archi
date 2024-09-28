package db

import (
	"log"

	"github.com/sakaguchi-0725/echo-onion-arch/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.DBConfig) (*gorm.DB, error) {
	dsn := config.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}
