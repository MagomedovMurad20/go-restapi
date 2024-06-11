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

// getAlbums returns all the albums as a JSON response.
func getAlbums(c *gin.Context) {
	// Return the albums with a status code of 200 (OK)
	// format the JSON response using indentation for better readability.
	c.IndentedJSON(http.StatusOK, albums)
}

// main initializes the Gin router and sets up the "/albums" route
// with the getAlbums handler function.

func main() {
	// Initialize Gin router with default middleware.
	router := gin.Default()

	// Set up the "/albums" route with the getAlbums handler function.
	// The getAlbums function returns all the albums as a JSON response.
	router.GET("/albums", getAlbums)
	// Run the server on the localhost at port 8080.
	router.Run("localhost:8080")
}
