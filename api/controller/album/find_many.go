package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian2s/web-service-gin/internal/infra/database"
)

func FindMany(c *gin.Context) {

	var album database.Album

	c.IndentedJSON(http.StatusOK, album.FindMany())
}
