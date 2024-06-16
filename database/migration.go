package database

import (
	"openanimalrescue-database/logging"
	"openanimalrescue-database/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Migration{},
		&models.User{},
		&models.Role{},
		&models.Animal{},
		&models.Venue{},
		&models.Adoption{},
		&models.PopUpDay{},
		&models.AnimalHealthStatus{},
		&models.AnimalSterilisationStatus{},
		&models.AnimalLocation{},
		&models.AdoptionStatus{},
	)
	if err != nil {
		logging.LogError("Failed to migrate database: " + err.Error())
		return
	}

	// Insert initial migration record
	db.Create(&models.Migration{MigrationID: "initial_migration"})
	logging.LogSuccess("Database migration completed successfully.")
}
