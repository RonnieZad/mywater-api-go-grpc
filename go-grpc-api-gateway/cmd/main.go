package main

import (
	"log"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/application"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/config"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment"
	
	"github.com/gin-gonic/gin"
)

func main() {

	// gin.SetMode(gin.ReleaseMode)

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	// Create a middleware to add the CORS headers
	corsMiddleware := func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}

	// Add the CORS middleware to the router
	r.Use(corsMiddleware)

	authSvc := *auth.RegisterRoutes(r, c)
	payment.RegisterRoutes(r, c, &authSvc)
	application.RegisterRoutes(r, c, &authSvc)

	r.Run(c.Port)
}

// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/application"
// 	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth"
// 	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/config"
// 	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment"
// 	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/property"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	c, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalln("Failed at config", err)
// 	}

// 	r := setupRouter(c)
// 	r.Run(c.Port)
// }

// func setupRouter(c config.Config) *gin.Engine {
// 	r := gin.Default()

// 	// Middleware: CORS
// 	r.Use(corsMiddleware)

// 	// Middleware: Error Handling
// 	r.Use(errorHandlingMiddleware)

// 	// Middleware: Authentication
// 	authSvc := auth.RegisterRoutes(r, c)

// 	// Modules
// 	property.RegisterRoutes(r, c, authSvc)
// 	payment.RegisterRoutes(r, c, authSvc)
// 	application.RegisterRoutes(r, c, authSvc)

// 	return r
// }

// func corsMiddleware(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 	if c.Request.Method == "OPTIONS" {
// 		c.AbortWithStatus(http.StatusOK)
// 		return
// 	}
// 	c.Next()
// }

// func errorHandlingMiddleware(c *gin.Context) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			// Handle panics, log the error, and return an appropriate response
// 			log.Printf("Panic occurred: %v", r)
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "An internal server error occurred.",
// 			})
// 		}

// 		if len(c.Errors) > 0 {
// 			// Handle errors from previous middleware or handlers
// 			// You can log the errors or return specific error responses
// 			for _, err := range c.Errors {
// 				log.Printf("Error occurred: %v", err)
// 			}
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Bad request. Please check your input.",
// 			})
// 		}
// 	}()

// 	// Continue with the next middleware or handler
// 	c.Next()
// }
