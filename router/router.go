package router

import (
	"example/web-service-gin/declarations"
	handler "example/web-service-gin/route_handlers"

	"github.com/gin-gonic/gin"
)

// albums: slice to seed record album data.
var albumsList = []declarations.Album{
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
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	},
}

func SetupGinRouter() *gin.Engine {
	// Engine instance with the Logger and Recovery middleware already attached.
	var router *gin.Engine = gin.Default()

	// GET is a shortcut for router.Handle("GET", path, handlers).

	router.GET("/albums", handler.GetAlbums((&albumsList)))

	router.GET("/albums/:id", handler.GetAlbumById(&albumsList))

	router.POST("/albums", handler.PostAlbums(&albumsList))

	router.GET("/", handler.HomeRoute)

	return router
}
