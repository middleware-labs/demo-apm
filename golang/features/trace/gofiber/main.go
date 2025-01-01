package main

import (
	"context"
	"errors"
	"log"


	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/contrib/otelfiber"
	"go.opentelemetry.io/otel"
	track "github.com/middleware-labs/golang-apm/tracker" 

)

var tracer = otel.Tracer("fiber-server")

func main() {
	config,_ := track.Track( 
		track.WithConfigTag("service", "your service name"),
		track.WithConfigTag("accessToken", "your access token"),
	)
	tp := config.Tp  
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	app := fiber.New()
	app.Use(otelfiber.Middleware())


	app.Get("/error", func(ctx *fiber.Ctx) error {
		return errors.New("abc")
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		name := getUser(id)
		return c.JSON(fiber.Map{"id": id, name: name})
	})

	log.Fatal(app.Listen(":3000"))
}

func getUser( id string) string {
	 if id == "123" {
		return "otelfiber tester"
	}
	return "unknown"
}