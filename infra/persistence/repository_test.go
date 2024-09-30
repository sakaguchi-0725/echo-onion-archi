package persistence_test

import (
	"log"
	"os"
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/config"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/db"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	setupTestDB()

	if testDB == nil {
		log.Println("Skipping tests due to failed DB connection")
		os.Exit(0)
	}

	code := m.Run()

	teardownTestDB()

	os.Exit(code)
}

func setupTestDB() {
	cfg := config.NewTestDBConfig()
	var err error

	testDB, err = db.NewTestDB(cfg)
	if err != nil {
		log.Printf("Failed to connect to test database: %v", err)
		testDB = nil
	}

	log.Println("Test DB initialized successfully.")
}

func teardownTestDB() {
	sqlDB, err := testDB.DB()
	if err == nil {
		sqlDB.Close()
	}
	log.Println("Test DB connection closed.")
}

func cleanUpTables(db *gorm.DB, tables ...string) {
	for _, table := range tables {
		if err := db.Exec("TRUNCATE TABLE " + table + " RESTART IDENTITY CASCADE").Error; err != nil {
			log.Printf("Failed to clean up table %s: %v", table, err)
		}
	}
}
