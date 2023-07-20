package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian2s/web-service-gin/internal/infra/database"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	var album database.Album

	album.Delete(id)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted"})
}
