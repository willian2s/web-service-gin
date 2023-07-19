package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: uuid.New().String(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: uuid.New().String(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: uuid.New().String(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: uuid.New().String(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", GetAlbums)

	router.Run(":8080")
}
