package configs

import (
	"fmt"
	"openanimalrescue-database/logging"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logging.LogError("Error loading .env file")
		return nil
	}

	dsn := os.Getenv("DATABASE_DSN")
	env := os.Getenv("ENV")

	if dsn == "" || env == "" {
		logging.LogError("DATABASE_DSN or ENV not set in .env file")
		return nil
	}

	dsn = addEnvToDSN(dsn, env)

	return &Config{
		DSN: dsn,
	}
}

func addEnvToDSN(dsn, env string) string {
	var dbName string
	fmt.Sscanf(dsn, "dbname=%s", &dbName)
	newDBName := fmt.Sprintf("%s-%s", dbName, env)
	newDSN := fmt.Sprintf("%s dbname=%s", dsn, newDBName)
	return newDSN
}
