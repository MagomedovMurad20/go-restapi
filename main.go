package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// handlers
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			c.BindJSON(a)
			mem := albums[:i+1]
			albums = append(albums[:i], a)
			albums = append(albums, mem...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// main
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums/:id", deleteAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
