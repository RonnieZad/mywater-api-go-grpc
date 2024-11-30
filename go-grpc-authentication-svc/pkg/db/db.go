package db

import (
	"log"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.UserClient{}, &models.PhoneNumberOTP{}, &models.EditAccountRequest{})
	// db.AutoMigrate(&models.Role{})
	// db.AutoMigrate(&models.Permission{})
	// db.AutoMigrate(&models.RolePermission{})

	return Handler{db}
}
