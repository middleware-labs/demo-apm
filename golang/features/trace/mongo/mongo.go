package main

import (
	"github.com/gin-gonic/gin"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	track "github.com/middleware-labs/golang-apm/tracker"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mw_mongo "github.com/middleware-labs/golang-apm-mongo/mongo"
)

func main() {
	opts := options.Client()
	opts.Monitor = mw_mongo.NewMonitor()
	opts.ApplyURI("mongodb://localhost:27017")
	r := gin.Default()
	config, _ := track.Track(
		track.WithConfigTag("service", "Your service name"),
		track.WithConfigTag("projectName", "Your project name"),
		track.WithConfigTag("accessToken", "Your access token"),
	)
	r.Use(g.Middleware(config))
	r.GET("/test", func(c *gin.Context) {
		ctx := c.Request.Context()
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			panic(err)
		}
		db := client.Database("example")
		inventory := db.Collection("inventory")

		_, err = inventory.InsertOne(ctx, bson.D{
			{Key: "item", Value: "canvas"},
			{Key: "qty", Value: 100},
			{Key: "attributes", Value: bson.A{"cotton"}},
			{Key: "size", Value: bson.D{
				{Key: "h", Value: 28},
				{Key: "w", Value: 35.5},
				{Key: "uom", Value: "cm"},
			}},
		})
		if err != nil {
			panic(err)
		}
	})
	_ = r.Run(":8082")
}
