package main

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	nodeexporterURL *string
)

// LivenessHandler always returns HTTP 200, use ReadinessHandler instead
func LivenessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "alive", "status": http.StatusOK})
}

// ReadinessHandler indicates HTTP 200 if the df-backend's nodeexporter serves metrics, otherwise HTTP 503
func ReadinessHandler(c *gin.Context) {
	resp, err := http.Get(*nodeexporterURL)
	if err == nil && resp.StatusCode == 200 {
		c.JSON(http.StatusOK, gin.H{"message": "ready", "status": http.StatusOK})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unavailable", "status": http.StatusServiceUnavailable})
	}
}

// NotFoundHandler to indicate that requested resource could not be found
func NotFoundHandler(c *gin.Context) {
	// log this event as it could be an attempt to break in...
	log.Infoln("Not found, requested URL path:", c.Request.URL.Path)
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "requested resource not found", "status": http.StatusNotFound})
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
	nodeexporterURL = flag.String("nodeexporterURL", "http://localhost:8080/metrics", "URL of the df-backend's nodeexporter")
	flag.Parse()

	router := SetupRouter()

	log.Infoln("df-frontend start...")
	defer log.Infoln("df-frontend shutdown!")

	// set port via PORT environment variable
	router.Run() // default port is 8080 (where the df-backend's nodeexporter is running on by default)
}
