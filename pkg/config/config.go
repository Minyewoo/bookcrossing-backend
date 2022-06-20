package config

import (
	"bookcrossing-backend/pkg/storage"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	storage.DBCredentials
	SigningKey    string
	ServerAddress string
}

func ReadConfig() *Config {
	godotenv.Load()
	return &Config{
		DBCredentials: storage.DBCredentials{
			DriverName:     os.Getenv("DriverName"),
			DataSourceName: os.Getenv("DataSourceName"),
		},
		SigningKey:    os.Getenv("SigningKey"),
		ServerAddress: os.Getenv("ServerAddress"),
	}
}
