package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    Port              string
    AuthSvcUrl        string
    PaymentSvcUrl     string
    ApplicationSvcUrl string
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
        AuthSvcUrl:        os.Getenv("AUTH_SVC_URL"),
        PaymentSvcUrl:     os.Getenv("PAYMENT_SVC_URL"),
        ApplicationSvcUrl: os.Getenv("APPLICATION_SVC_URL"),
    }

    return config, nil
}
