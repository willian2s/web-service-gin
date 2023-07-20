package main

import (
	"github.com/gin-gonic/gin"

	albumController "github.com/willian2s/web-service-gin/api/controller/album"
)

func Router() *gin.Engine {
	router := gin.Default()

	album := router.Group("/albums")
	{
		album.POST("", albumController.CreateAlbum)
		album.GET("", albumController.FindMany)
		album.GET("/:id", albumController.FindOne)
		album.DELETE("/:id", albumController.Delete)
	}

	return router
}

func main() {
	router := Router()

	router.Run(":8080")
}
