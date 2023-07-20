package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	controller "github.com/willian2s/web-service-gin/api/controller/album"
)

type Routes struct {
	Route *gin.Engine
}

func NewRoutes(route *gin.Engine) *Routes {
	return &Routes{
		Route: route,
	}
}

func Execute() {
	port := os.Getenv("PORT")

	router := gin.Default()
	router.SetTrustedProxies(nil)

	album := router.Group("/albums")
	{
		album.POST("", controller.CreateAlbum)
		album.GET("", controller.FindMany)
		album.GET("/:id", controller.FindOne)
		album.DELETE("/:id", controller.Delete)
	}

	router.Run(fmt.Sprintf(":%s", port))
}
