package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willian2s/web-service-gin/internal/infra/database"
	"github.com/willian2s/web-service-gin/internal/usecase"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

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

	err = usecase.Delete(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}
