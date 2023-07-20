package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian2s/web-service-gin/internal/infra/database"
	"github.com/willian2s/web-service-gin/internal/usecase"
)

func CreateAlbum(c *gin.Context) {
	var newAlbum *usecase.AlbumInputDTO

	db, err := database.Execute()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	repository := database.NewAlbumRepository(db)
	usecase := usecase.AlbumUsecase{
		AlbumRepository: repository,
	}

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	album, err := usecase.Save(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, album)
}
