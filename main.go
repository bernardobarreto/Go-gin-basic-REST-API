package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:3000")
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "The flood", Artist: "Of Mice and Men", Price: 19.45},
	{ID: "2", Title: "Infest", Artist: "Papa Roach", Price: 20.50},
	{ID: "3", Title: "Meteora", Artist: "Linkin Park", Price: 22.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
