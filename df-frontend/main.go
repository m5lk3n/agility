package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// LivenessHandler TODO
func LivenessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"alive": true})
}

// ReadinessHandler TODO
func ReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ready": true})
}

// NotFoundHandler to indicate that requested resource could not be found
func NotFoundHandler(c *gin.Context) {
	// log this event as it could be an attempt to break in...
	log.Infoln("Not found, requested URL path:", c.Request.URL.Path)
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "requested resource not found"})
}

// SetupRouter is published here to allow setup of tests
func SetupRouter() *gin.Engine {

	router := gin.Default()

	// to debug: router.Use(gindump.Dump())

	router.Use(gin.Recovery()) // "recover from any panics", write 500 if any

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.NoRoute(NotFoundHandler)

	// public, generic API
	router.GET("/healthy", LivenessHandler)
	router.GET("/ready", ReadinessHandler)

	return router
}

func main() {
	router := SetupRouter()

	log.Infoln("web app start...")
	defer log.Infoln("web app shutdown!")

	// set port via PORT environment variable
	router.Run() // default port is 8080
}
