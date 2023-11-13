package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album: data about a record album | struct tags: for JSON serialization
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumNotFoundErr struct {
	Err string `json:"error"`
}

// albums: slice to seed record album data.
var albumsList = []Album{
	{
		ID:     "1001",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "1002",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
	{
		ID:     "1003",
		Title:  "Sarah Vaughan and CLifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	},
}

func main() {
	var ginRouter *gin.Engine = setupGinRouter()

	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	// It is a shortcut for http.ListenAndServe(addr, router)
	ginRouter.Run("localhost:8080")
}

func getAlbums(albums *[]Album) func(c *gin.Context) {
	return func(c *gin.Context) {

		// marshel Go object and write to response body | set content-type to application/json
		c.IndentedJSON(http.StatusOK, albums)
	}
}

func homeRoute(c *gin.Context) {

	// String writes the given string into the response body.

	c.String(http.StatusOK, "First Gin Web Server!")
}

func postAlbums(albums *[]Album) func(c *gin.Context) {
	return func(c *gin.Context) {

		var newAlbum Album

		// unmarshel data into GO object
		if err := c.BindJSON(&newAlbum); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// map[string]string{"error": err.Error()}

		*albums = append(*albums, newAlbum)

		log.Println("New Album Added:", newAlbum)

		c.IndentedJSON(http.StatusOK, albums)

	}

}

func getAlbumById(albums *[]Album) func(c *gin.Context) {
	return func(c *gin.Context) {

		// Param returns the value of the URL param.
		id := c.Param("id")

		for _, album := range *albums {
			if album.ID == id {
				c.IndentedJSON(http.StatusOK, album)
				return
			}

		}
		c.IndentedJSON(http.StatusOK, AlbumNotFoundErr{
			Err: fmt.Sprintf("could not locate an album with the id: %s", id),
		})

	}
}

func setupGinRouter() *gin.Engine {
	// Engine instance with the Logger and Recovery middleware already attached.
	var ginRouter *gin.Engine = gin.Default()

	// GET is a shortcut for router.Handle("GET", path, handlers).
	ginRouter.GET("/albums", getAlbums(&albumsList))
	ginRouter.GET("/albums/:id", getAlbumById(&albumsList))
	ginRouter.POST("/albums", postAlbums(&albumsList))

	// GET is a shortcut for router.Handle("GET", path, handlers).
	ginRouter.GET("/", homeRoute)

	return ginRouter
}

// Developing A RESTful API with GO and Gin Web Framework

//   - write a RESTFul Web Service API
