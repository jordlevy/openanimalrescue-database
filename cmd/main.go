package main

import (
	"openanimalrescue-database/configs"
	"openanimalrescue-database/database"
	"openanimalrescue-database/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := configs.GetConfig()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		logging.LogError("Failed to connect database: " + err.Error())
		return
	}

	database.Migrate(db)
	database.Seed(db)
}
