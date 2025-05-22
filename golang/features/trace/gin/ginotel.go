package main

import (
	"net/http"
	// "errors"
	"github.com/gin-gonic/gin"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	"github.com/middleware-labs/golang-apm/logger"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	r := gin.Default()
	config, _ := track.Track(
		track.WithConfigTag("service", "golang-gin"),
		track.WithConfigTag("accessToken", "lneevcqdcnzccpdpbhelkxpjffdxeznkryvb"),
		track.WithConfigTag("target", "ruplp.middleware.io:443"),
		track.WithConfigTag(track.Debug,true),
		// track.WithConfigTag(track.PauseMetrics,true),
	)
	// logs
	logger.Error("Error")
	logger.Info("Info")
	logger.Warn("Warn")

	r.Use(g.Middleware(config))
	r.GET("/books", FindBooks)
	r.Run(":8090")
}

func FindBooks(c *gin.Context) {
    // Simulate a panic to generate an error
    panic("failed to fetch books from database")

    // The following code won't be executed due to the panic
    // but is left here for clarity
    c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error occurred"})
}
