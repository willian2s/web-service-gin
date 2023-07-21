package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian2s/web-service-gin/internal/infra/database"
	"github.com/willian2s/web-service-gin/internal/usecase"
)

func FindMany(c *gin.Context) {
	db, err := database.Execute()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	defer db.Close()

	repository := database.NewAlbumRepository(db)
	usecase := usecase.AlbumUsecase{
		AlbumRepository: repository,
	}

	albums, err := usecase.FindMany()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}
