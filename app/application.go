package app

import (
	"github.com/JayCardenas/PhotoGallary/controllers/photos"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	// // method for routing
	// urlMapping()

	// GET Image
	router.GET("/", photos.Get)

	// POST Image
	router.POST("/", photos.Create)

	// Starts listening and servicing HTTP requests on port 8000
	router.Run(":8000")

}
