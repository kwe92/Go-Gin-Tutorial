package route_handlers

import (
	"example/web-service-gin/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeRoute(c *gin.Context) {

	// String writes the given string into the response body.
	c.String(http.StatusOK, "First Gin Web Server!")
}

func GetAlbums(albums *[]model.Album) func(c *gin.Context) {
	return func(c *gin.Context) {

		// marshel Go object and write to response body | set content-type to application/json
		c.IndentedJSON(http.StatusOK, albums)
	}
}

func PostAlbums(albums *[]model.Album) func(c *gin.Context) {
	return func(c *gin.Context) {

		var newAlbum model.Album

		// unmarshel data into GO object
		if err := c.ShouldBindJSON(&newAlbum); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		*albums = append(*albums, newAlbum)

		log.Println("New model.Album Added:", newAlbum)

		c.IndentedJSON(http.StatusOK, albums)

	}

}

func GetAlbumById(albums *[]model.Album) func(c *gin.Context) {
	return func(c *gin.Context) {

		// Param returns the value of the URL param.
		id := c.Param("id")

		for _, album := range *albums {
			if album.ID == id {
				c.IndentedJSON(http.StatusOK, album)
				return
			}

		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("could not locate an album with the id: %s", id)})

	}
}

// gin.Context | struct | gin package

//   - a struct that is the most important part of gin
//   - allows you to:

//       ~ pass variables between middleware
//       ~ manage data flow
//       ~ validate JSON of a request
//       ~ serialize JSON
