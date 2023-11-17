package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	"github.com/middleware-labs/golang-apm/logger"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	config, _ := track.Track(
		track.WithConfigTag("service", "Your service Name"),
		track.WithConfigTag("projectName", "your project name"),
	)
	// logs
    logger.Error("Error")
    logger.Info("Info")
    logger.Warn("Warn")

	r := gin.Default()
	r.Use(g.Middleware(config))
	r.GET("/books", FindBooks)
	r.Run(":8090")
}

func FindBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
