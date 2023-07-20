package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian2s/web-service-gin/internal/infra/database"
	usecase "github.com/willian2s/web-service-gin/internal/usecase/album"
)

func CreateAlbum(c *gin.Context) {
	var newAlbum *usecase.AlbumInputDTO

	db, err := database.Execute()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	repository := database.NewAlbumRepository(db)
	usecase := usecase.CreateAlbum{
		AlbumRepository: repository,
	}

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	album, err := usecase.Handle(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, album)
}
