package config

import "os"

type (
	AppConfig struct {
		PORT             string
		PostgresHost     string
		PostgresPort     string
		PostgresDBName   string
		PostgresUsername string
		PostgresPassword string
	}
)

var (
	config *AppConfig = nil
)

func GetConfig() *AppConfig {
	if config != nil {
		return config
	}

	config = &AppConfig{
		PORT:             os.Getenv("PORT"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDBName:   os.Getenv("POSTGRES_DB_NAME"),
		PostgresUsername: os.Getenv("POSTGRES_USERNAME"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
	}

	return config
}
