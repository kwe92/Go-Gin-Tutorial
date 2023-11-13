package main

import (
	"example/web-service-gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	var ginRouter *gin.Engine = router.SetupGinRouter()

	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	// It is a shortcut for http.ListenAndServe(addr, router)
	ginRouter.Run("localhost:8080")
}

// Develop RESTful API with GO and Gin Web Framework

//   - write a RESTFul Web Service API
