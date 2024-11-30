package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	DBUrl             string
	JWTSecretKey      string
	ApplicationSvcUrl string
	PaymentSvcUrl     string
	PropertySvcUrl    string
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
		Port:              os.Getenv("PORT"),
		DBUrl:             os.Getenv("DB_URL"),
		JWTSecretKey:      os.Getenv("JWT_SECRET_KEY"),
		ApplicationSvcUrl: os.Getenv("APPLICATION_SVC_URL"),
		PaymentSvcUrl:     os.Getenv("PAYMENT_SVC_URL"),
		PropertySvcUrl:    os.Getenv("PROPERTY_SVC_URL"),
	}

	return config, nil
}
