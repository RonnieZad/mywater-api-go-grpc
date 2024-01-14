package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	DBUrl          string `mapstructure:"DB_URL"`
	AuthSvcUrl     string `mapstructure:"AUTH_SVC_URL"`
	PropertySvcUrl string `mapstructure:"PROPERTY_SVC_URL"`
}

func LoadConfig() (*Config, error) {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return nil, err
	}

	// Create a new Config object with the environment variables
	config := &Config{
		Port:           os.Getenv("PORT"),
		DBUrl:          os.Getenv("DB_URL"),
		AuthSvcUrl:     os.Getenv("AUTH_SVC_URL"),
		PropertySvcUrl: os.Getenv("PROPERTY_SVC_URL"),
	}

	return config, nil
}
