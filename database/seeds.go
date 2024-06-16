package database

import (
	"openanimalrescue-database/logging"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Add initial data here if needed
	logging.LogInfo("Database seeding completed.")
}
