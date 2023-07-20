package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian2s/web-service-gin/internal/infra/database"
)

func CreateAlbum(c *gin.Context) {
	var newAlbum database.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.Save()

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
