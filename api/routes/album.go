package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/willian2s/web-service-gin/api/controller/album"
)

func AlbumRoute(router *gin.Engine) {
	album := router.Group("/albums")
	{
		album.POST("", controller.CreateAlbum)
		album.GET("", controller.FindMany)
		album.GET("/:id", controller.FindOne)
		album.DELETE("/:id", controller.Delete)
	}
}
