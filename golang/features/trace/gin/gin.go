package main

import (
	"github.com/gin-gonic/gin"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	track "github.com/middleware-labs/golang-apm/tracker"
	"net/http"
)

func main() {
	r := gin.Default()
	config, _ := track.Track(
		track.WithConfigTag("service", "your service name"),
		track.WithConfigTag("accessToken", "your access token"),
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
	ctx := c.Request.Context()
	track.SetAttribute(ctx, "user.id", "1")
	track.SetAttribute(ctx, "user.role", "admin")
	track.SetAttribute(ctx, "user.scope", "read:message,write:files")
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
