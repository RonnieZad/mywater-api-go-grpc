// Package db provides the database handler for the application.
package db

import (
	"log"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Handler struct contains the DB object from the gorm library.
type Handler struct {
	DB *gorm.DB
}

// Init function initializes a new database connection using the provided url.
// It also sets up the database schema by auto migrating the models.
// It returns a Handler struct containing the DB object.
func Init(url string) Handler {
	// Open a new db connection.
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	// If there's an error, log it and stop the execution.
	if err != nil {
		log.Fatalln(err)
	}

	// Auto migrate the models.
	db.AutoMigrate(&models.Advert{})
	db.AutoMigrate(&models.LabelScan{})

	// Return the handler.
	return Handler{db}
}
