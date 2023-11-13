package router

import (
	"example/web-service-gin/declarations"
	handlers "example/web-service-gin/route_handlers"

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
		Title:  "Sarah Vaughan and CLifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	},
}

func SetupGinRouter() *gin.Engine {
	// Engine instance with the Logger and Recovery middleware already attached.
	var ginRouter *gin.Engine = gin.Default()

	// GET is a shortcut for router.Handle("GET", path, handlers).
	ginRouter.GET("/albums", handlers.GetAlbums((&albumsList)))

	ginRouter.GET("/albums/:id", handlers.GetAlbumById(&albumsList))

	ginRouter.POST("/albums", handlers.PostAlbums(&albumsList))

	ginRouter.GET("/", handlers.HomeRoute)

	return ginRouter
}
