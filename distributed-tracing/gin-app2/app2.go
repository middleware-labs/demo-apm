package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	track "github.com/middleware-labs/golang-apm/tracker"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	config, _ := track.Track(
		track.WithConfigTag("service", "go-app-b"),
		track.WithConfigTag("projectName", "go-app-b"),
	)
	tp := config.Tp
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	r := gin.New()
	r.Use(otelgin.Middleware(config.ServiceName, otelgin.WithTracerProvider(tp)))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost},
	}))
	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Printf("Debug HEADER:%v", c.Request.Header)
		fmt.Println("ID:", id)
		c.JSON(200, id)
	})
	_ = r.Run(":3001")
}
